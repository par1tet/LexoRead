from jwt import encode
from config import JWT_SECRET


def generate_jwt(id, username, email, password):
    return encode(
        payload={
            "id": id,
            "username": username,
            "email": email,
            "password": password
        },
        key=JWT_SECRET
    )