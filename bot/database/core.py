from sqlalchemy import select
from database.database import async_engine, async_session, Base
from database.models import Admins, Questions


async def create_tables():
    async with async_engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def create_question(text, chat_id):
    async with async_session() as session:
        question = Questions(text=text, chat_id=chat_id)
        session.add(question)
        await session.commit()


async def get_questions(page):
    async with async_session() as session:
        questions = select(Questions).where(Questions.admin==None).limit(8).offset(page * 8 + 1)
        result = await session.execute(questions)
        return [str(i[0].id) for i in result.all()]


async def leave_from_question(chat_id):
    async with async_session as session:
        admin = await session.execute(select(Admins).where(Admins.chat_id==chat_id)).one()[0]
        admin.question = None
        question = await session.execute(select(Questions).where(Questions.chat_id==admin.question.chat_id)).one()[0]
        question.admin = None
        return admin, question


async def get_question(id):
    async with async_session() as session:
        question = await session.execute(select(Questions).where(Questions.id==int(id)))
        data = question.one()[0]
        return data


async def get_admins():
    async with async_session() as session:
        result = await session.execute(select(Admins))
        answer = [i[0].chat_id for i in result.all()]
        return answer


async def get_chat_id_questioner(chat_id):
    async with async_session() as session:
        admin = await session.execute(select(Admins).where(Admins.chat_id==chat_id))
        question = await session.execute(select(Questions).where(Questions.admin==admin.id))
        return question.one()[0].chat_id


async def start_connection(question_id, admin_chat_id):
    async with async_session() as session:
        admin = await session.execute(select(Admins).where(Admins.chat_id==admin_chat_id))
        admin.quesion = question_id
        question = await session.get(Questions, question_id)
        question.admin = admin.id
        await session.commit()
