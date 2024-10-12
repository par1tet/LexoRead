from sqlalchemy import Column, Integer, String, ForeignKey
from sqlalchemy.orm import mapped_column, Mapped, relationship
from database.database import Base


class Questions(Base):
    __tablename__ = "questions"

    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)

    text: Mapped[str]

    chat_id: Mapped[int]

    admin: Mapped["Admins" | None] = relationship(back_populates="question")


class Admins(Base):
    __tablename__ = "admins"

    id = Column(Integer, primary_key=True, autoincrement=True)

    chat_id = Column(Integer)

    quesiton: Mapped["Questions" | None] = relationship(back_populates="admin")