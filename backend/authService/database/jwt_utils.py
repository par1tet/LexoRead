from jwt import encode


def generate_jwt(id, username, email, password):
    return encode({
        "id": id,
        "username": username,
        "email": email,
        "password": password
    })
