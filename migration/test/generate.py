import csv
from datetime import datetime
from hashlib import md5
from random import randint
from typing import Sequence

hash = lambda x: md5(x.encode()).hexdigest()


def gen_users(count: int) -> Sequence[dict]:
    users = []
    for i in range(1, count + 1):
        user = {
            "id": i,
            "login": str(i),
            "password": hash(str(i)),
            "registration_time": datetime(2023, 3, i % 22).isoformat(),
            "nickname": str(i),
        }
        users.append(user)
    return users


def gen_sessions(count: int) -> Sequence[dict]:
    sessions = []

    for i in range(1, count + 1):
        session = {
            "token": hash(str(i)),
            "user_id": i,
            "expired_at": datetime(2023, 3, i % 22).isoformat(),
        }
        sessions.append(session)

    return sessions


def gen_servers(count: int) -> Sequence[dict]:
    servers = []

    for i in range(1, count + 1):
        server = {
            "id": i,
            "name": hash(str(i)),
            "creation_time": datetime(2023, 3, i % 22).isoformat(),
            "owner_id": i,
        }
        servers.append(server)
    return servers


def gen_server_profiles(count: int) -> Sequence[dict]:
    return [
        {
            "server_id": i,
            "user_id": i,
            "join_time": datetime(2023, 3, i % 22).isoformat(),
            "nickname": str(hash(str(i))),
        }
        for i in range(1, count + 1)
    ]


def gen_channels(count: int) -> Sequence[dict]:
    return [
        {
            "id": i,
            "name": hash(str(i)),
            "creation_time": datetime(2023, 3, i % 22).isoformat(),
            "server_id": i,
            "creator_id": i,
        }
        for i in range(1, count + 1)
    ]


def gen_roles(count: int) -> Sequence[dict]:
    return [
        {
            "id": i,
            "name": hash(str(i)),
            "permission": randint(0, 32),
            "creation_time": datetime(2023, 3, i % 22).isoformat(),
            "server_id": i,
            "created_by": i,
        }
        for i in range(1, count + 1)
    ]


def gen_permissions(count: int) -> Sequence[dict]:
    return [
        {
            "value": i,
            "name": str(i),
            "description": hash(str(i)),
        }
        for i in range(1, count + 1)
    ]


def automatic_shit(name: str, count: int):
    func = globals()[f"gen_{name}s"]
    with open(f"{name}.csv", "w") as f:
        fields = list(func(1)[0].keys())
        w = csv.DictWriter(f, fields)
        w.writeheader()
        w.writerows(func(count))


names = [
    "user",
    "session",
    "server",
    "server_profile",
    "channel",
    "role",
    "permission",
]


def gen_db_dataset():
    count = 10

    for name in names:
        automatic_shit(name, count)


def gen_copy_file():
    with open("load_test_data.psql", "w") as f:
        for name in names:
            c = f'\copy "{name}" from \'migration/test/{name}.csv\' csv header'
            f.write(f"psql -U postgres -d openchat -c \'{c}\'\n")


if __name__ == "__main__":
    gen_db_dataset()
    gen_copy_file()
