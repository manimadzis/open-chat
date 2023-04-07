CREATE TABLE IF NOT EXISTS "user"
(
    id                BIGSERIAL PRIMARY KEY,
    login             TEXT      NOT NULL UNIQUE,
    password          TEXT      NOT NULL,
    registration_time TIMESTAMP NOT NULL,
    nickname          TEXT      NOT NULL
);


CREATE TABLE IF NOT EXISTS session
(
    token      TEXT PRIMARY KEY,
    user_id    BIGINT UNIQUE NOT NULL REFERENCES "user" (id),
    expired_at TIMESTAMP     NOT NULL
);

CREATE TABLE IF NOT EXISTS server
(
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT      NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    owner_id      BIGINT REFERENCES "user" (id)
);


CREATE TABLE IF NOT EXISTS server_profile
(
    server_id BIGINT    NOT NULL REFERENCES server (id),
    user_id   BIGINT    NOT NULL REFERENCES "user" (id),
    join_time TIMESTAMP NOT NULL,
    nickname  TEXT      NOT NULL,
    PRIMARY KEY (server_id, user_id)
);


CREATE TABLE IF NOT EXISTS channel
(
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT      NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    server_id     BIGINT REFERENCES server (id),
    creator_id    BIGINT REFERENCES "user" (id)
);


CREATE TABLE IF NOT EXISTS role
(
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT      NOT NULL,
    permission    BIGINT    NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    server_id     BIGINT REFERENCES server (id),
    created_by    BIGINT REFERENCES "user" (id)
);


CREATE TABLE IF NOT EXISTS permission
(
    value       BIGINT PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sticker_pack
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sticker
(
    id              BIGSERIAL PRIMARY KEY,
    path            TEXT NOT NULL,
    sticker_pack_id BIGINT REFERENCES sticker_pack (id)
);


CREATE TABLE IF NOT EXISTS message
(
    id            BIGSERIAL PRIMARY KEY,
    text          TEXT      NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    sender_id     BIGINT REFERENCES "user" (id),
    channel_id    BIGINT REFERENCES channel (id),
    sticker_id    BIGINT REFERENCES sticker (id)
);


