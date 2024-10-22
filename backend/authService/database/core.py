from sqlalchemy import select
from database.database import async_engine, async_session, Base
from database.models import User


async def create_tables():
    async with async_engine() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def reg_user(username, email, password):
    async with async_session() as session:
        user = User(username=username, email=email, password=password)
        session.add(user)
        await session.commit()


async def log_user(username, password):
    async with async_session() as session:
        user = await session.execute(select(User).where(User.password==password and User.username==username))
        return user.scalar_one()