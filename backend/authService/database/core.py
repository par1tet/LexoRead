from sqlalchemy import select
from database.database import async_engine, async_session, Base
from database.models import Users


async def create_tables():
    async with async_engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def reg_user(username, email, password):
    async with async_session() as session:
        user = Users(username=username, email=email, password=password)
        session.add(user)
        await session.commit()
        return user


async def log_user(username, password):
    async with async_session() as session:
        user = await session.execute(select(Users).where(Users.username==username, Users.password==password))
        return user.scalar_one_or_none()