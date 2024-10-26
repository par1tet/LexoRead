from jwt import encode, decode
from config import JWT_SECRET
from database.core import log_user


def generate_jwt(id, username, email, password):
    return encode(
        payload={
            "id": id,
            "username": username,
            "email": email,
            "password": password
        },
        key=JWT_SECRET,
        algorithm="HS256"
    )


def decode_jwt(jwt_string):
    decoded = decode(
        jwt=jwt_string,
        key=JWT_SECRET,
        algorithms=["HS256"]
    )
    return decoded