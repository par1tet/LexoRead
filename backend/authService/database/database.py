from sqlalchemy.ext.asyncio import create_async_engine, async_sessionmaker
from sqlalchemy.orm import DeclarativeBase
from config import DATABASE_URL

async_engine = create_async_engine(
    url=DATABASE_URL,
)

async_session = async_sessionmaker(async_engine)

class Base(DeclarativeBase):
    pass