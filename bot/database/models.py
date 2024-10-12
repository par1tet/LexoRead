from typing import Optional
from sqlalchemy import ForeignKey, BigInteger, Column
from sqlalchemy.orm import mapped_column, Mapped
from database.database import Base


class Questions(Base):
    __tablename__ = "questions"

    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)

    text: Mapped[str]

    chat_id = Column(BigInteger)

    admin_id: Mapped[Optional[int]] = mapped_column(ForeignKey("admins.id"))


class Admins(Base):
    __tablename__ = "admins"

    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)

    chat_id = Column(BigInteger)

    quesiton: Mapped[Optional[int]] = mapped_column(ForeignKey("questions.id"))