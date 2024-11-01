from sqlalchemy import Integer, String, Boolean, BIGINT, ARRAY, TEXT
from sqlalchemy.orm import mapped_column
from database.database import Base


class Users(Base):
    __tablename__ = "users"

    id = mapped_column(Integer, primary_key=True, autoincrement=True)
    isBanned = mapped_column(Boolean)
    username = mapped_column(String(256))
    email = mapped_column(String(256))
    description = mapped_column(String)
    hashPassword = mapped_column(TEXT)
    favBooks = mapped_column(ARRAY(BIGINT), default=[])