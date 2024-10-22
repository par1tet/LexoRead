import asyncio

from fastapi import FastAPI
from uvicorn import Server, Config
from database.core import create_tables, log_user, reg_user
from database.jwt_utils import generate_jwt


app = FastAPI()


@app.post("/reg")
async def registration(username, email, password):
    user = await reg_user(username, email, password)
    return generate_jwt(user.id, username, email, password)


@app.post("/log")
async def login(username, password):
    if log_user():
        return 
    return


async def main():
    await create_tables()
    server = Server(Config(app))
    await server.serve()


if __name__ == "__main__":
    asyncio.run(main())