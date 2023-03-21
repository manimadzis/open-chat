CREATE TABLE IF NOT EXISTS session(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL,
    token TEXT UNIQUE NOT NULL,
    expired_at TIMESTAMP NOT NULL
--     FOREIGN KEY user_id REFERENCES user(id)
);


CREATE TABLE IF NOT EXISTS "user"(
    id BIGSERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    registration_time  TIMESTAMP NOT NULL,
    nickname TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS role(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    permission BIGINT,
    creation_time TIMESTAMP NOT NULL,
    server_id BIGINT REFERENCES server(id),
    created_by BIGINT REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS server(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id BIGINT REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS server_profile(
    server_id BIGINT NOT NULL REFERENCES server(id),
    user_id BIGINT NOT NULL REFERENCES "user"(id),
    join_time TIMESTAMP NOT NULL,
    PRIMARY KEY (server_id, user_id)
);

CREATE TABLE IF NOT EXISTS server_profile(

)