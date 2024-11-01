from sqlalchemy import select
from database.database import async_engine, async_session, Base
from database.models import Users
from bcrypt import hashpw, gensalt, checkpw
from database.jwt_utils import generate_jwt

async def create_tables():
    async with async_engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def reg_user(username, email, password):
    async with async_session() as session:
        user = Users(
            isBanned=False,
            username=username,
            email=email,
            description="",
            hashPassword=hashpw(password.encode("utf-8"), gensalt(6)).decode()
        )
        jwt_string = generate_jwt(user.id, user.username, user.email, user.hashPassword)
        session.add(user)
        await session.commit()
        return jwt_string


async def log_user(username, password):
    async with async_session() as session:
        selected = await session.execute(
            select(Users).where(
                Users.username==username
            )
        )
        user = selected.scalar_one_or_none()
        if user is None:
            return False

        if checkpw(password.encode(), ):
            return generate_jwt(user.id, username, user.email, user.hashPassword)

        return False