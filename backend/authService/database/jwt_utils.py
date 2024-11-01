from jwt import encode, decode
from config import JWT_SECRET


def generate_jwt(id, username, email, hashPassword):
    return encode(
        payload={
            "id": id,
            "username": username,
            "email": email,
            "hashPassword": hashPassword
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