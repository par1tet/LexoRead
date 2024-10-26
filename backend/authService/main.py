import asyncio

from fastapi import FastAPI
from uvicorn import Server, Config
from database.core import create_tables, log_user, reg_user
from database.jwt_utils import generate_jwt, decode_jwt
from models import User, UserWithEmail, JWTSecret

app = FastAPI()


@app.post("/reg")
async def registration(user: UserWithEmail):
    user = await reg_user(user.username, user.email, user.password)
    return generate_jwt(user.id, user.username, user.email, user.password)


@app.post("/log")
async def login(user: User):
    user = await log_user(user.username, user.password)
    if user:
        return generate_jwt(user.id, user.username, user.email, user.password)
    return False


@app.post("/is_auth")
async def is_auth(jwt: JWTSecret):
    decoded = decode_jwt(jwt.jwt_string)
    username, password = decoded["username"], decoded["password"]
    return not (await log_user(username, password)) is None


async def main():
    await create_tables()
    server = Server(Config(app))
    await server.serve()


if __name__ == "__main__":
    asyncio.run(main())