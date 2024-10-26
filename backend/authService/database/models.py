from sqlalchemy import Integer, String
from sqlalchemy.orm import mapped_column
from database.database import Base


class Users(Base):
    __tablename__ = "users"

    id = mapped_column(Integer, primary_key=True, autoincrement=True)
    username = mapped_column(String(15))
    email = mapped_column(String)
    password = mapped_column(String)