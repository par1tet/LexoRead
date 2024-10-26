from database.models import Users
from pydantic import BaseModel

class User(BaseModel):
    username: str
    password: str


class UserWithEmail(User):
    email: str


class JWTSecret(BaseModel):
    jwt_string: str