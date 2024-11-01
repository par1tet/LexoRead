import asyncio

from fastapi import FastAPI
from uvicorn import Server, Config
from database.core import create_tables, log_user, reg_user
from database.jwt_utils import generate_jwt, decode_jwt
from models import User, UserForReg, JWTSecret


app = FastAPI()


@app.post("/reg")
async def registration(user: UserForReg):
    return await reg_user(user.username, user.email, user.password)


@app.post("/log")
async def login(user: User):
    jwt_string = await log_user(user.username, user.password)
    return jwt_string or False


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