from sqlalchemy import ForeignKey, BIGINT, Column
from sqlalchemy.orm import mapped_column, Mapped
from database.database import Base


class Questions(Base):
    __tablename__ = "questions"

    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)

    text: Mapped[str]

    chat_id = Column(BIGINT)

    admin: Mapped[BIGINT] = mapped_column(ForeignKey("admins.chat_id"))


class Admins(Base):
    __tablename__ = "admins"

    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)

    chat_id = Column(BIGINT)

    quesiton: Mapped[BIGINT] = mapped_column(ForeignKey("questions.chat_id"))