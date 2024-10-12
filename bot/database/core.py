from sqlalchemy import select, text
from database.database import async_engine, async_session, Base
from database.models import Admins, Questions


async def create_tables():
    async with async_engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def create_question(text, chat_id):
    async with async_session() as session:
        question = Questions(text=text, chat_id=chat_id)
        session.add(question)
        session.commit()


async def get_questions():
    async with async_session() as session:
        questions = select(Questions)
        result = await session.execute(questions)
        print(result.one()[0].id)
        answer = []
        for question in list(result.one()):
            answer.append(str(question.id))
        return answer


async def leave_from_question(chat_id):
    async with async_session as session:
        admin = select(Admins).where(Admins.chat_id==chat_id).as_scalar().all_()[0]
        admin.question = None
        question = select(Questions).where(Questions.chat_id==admin.question.chat_id).as_scalar().all_()[0]
        question.admin = None
        return admin, question