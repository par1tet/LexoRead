from fastapi import FastAPI
from database.core import reg_user

app = FastAPI()


@app.post("/reg")
async def registration(username, password):
    reg_user(username, password)