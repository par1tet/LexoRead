from fastapi import FastAPI
from database.core import reg_user, log_user

app = FastAPI()


@app.post("/reg")
async def registration(username, password):
    reg_user(username, password)


@app.post("/log")
async def login(username, password):
    if log_user():
        return
    return