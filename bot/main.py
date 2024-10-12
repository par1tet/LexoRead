import logging
import asyncio

from aiogram import Bot, Dispatcher, types
from aiogram.filters import Command, CommandStart, CommandObject, BaseFilter

from config import TOKEN
from database.core import create_tables, create_question, get_questions, leave_from_question


class AdminFilter(BaseFilter):
    def __init__(self):
        self.admins = [1340572920]

    async def __call__(self, message: types.CallbackQuery) -> bool:
        admins = self.admins.copy()
        if isinstance(message, types.Message) and message.chat.id in admins: # TODO: add getting admin list from db
            return True
        if message.message.chat.id in admins:
            return True
        await message.answer("Нет прав.")
        return False


class ConnectionFilter(BaseFilter):
    def __init__(self, reverse: bool = False):
        self.reverse = reverse

    async def __call__(self, message: types.Message):
        if not self.reverse: 
            if True: # TODO: add check exists connectiong
                await message.answer("Вы не в чате с вопросом.")
                return False
            return True
        else:
            if False: # TODO: add check exists connectiong
                await message.answer("Вы уже в чате с вопросом.")
                return False
            return True


logging.basicConfig(level=logging.INFO)
bot = Bot(TOKEN)
dp = Dispatcher()


questions = [
    {"question": "How i can create the book?", "chat_id": 1340572920}
]


@dp.callback_query(AdminFilter())
async def btn_handler(call: types.CallbackQuery):
    data = questions[int(call.data)]
    question = data["question"]
    text = "Нажмите чтобы ответить:"

    kb = [[types.InlineKeyboardButton(
        text="Ответить", url=f"t.me/testttstesbot?start={data["chat_id"]}"
    )]]

    markup = types.InlineKeyboardMarkup(inline_keyboard=kb)

    await call.message.answer(question + "\n\n" + text, reply_markup=markup)


@dp.message(CommandStart(), AdminFilter(), ConnectionFilter(True))
async def start_answering(message: types.Message, command: CommandObject):
    # TODO: connect with questioner
    await message.answer("Вы начали диалог с вопросом.")


@dp.message(CommandStart())
async def start(message: types.Message):
    await message.answer("Отправьте свой вопрос тех. поддержке")


@dp.message(Command("admin"), AdminFilter())
async def admin(message: types.Message):
    await message.answer(
"""
Просмотреть вопросы:
/check
Закрыть вопрос:
/cancel
Выйти из вопроса:
/leave
"""
    )


@dp.message(Command("cancel"), AdminFilter(), ConnectionFilter())
async def cancel_answering(message: types.Message):
    # TODO: stop connection with questioner and delete question
    await message.answer("Вы покинули чат с вопросом и удалили его.")


@dp.message(Command("leave"), AdminFilter(), ConnectionFilter())
async def cancel_answering(message: types.Message):
    print(await leave_from_question(message.chat_id))
    await message.answer("Вы покинули чат с вопросом.")


@dp.message(Command("check"), AdminFilter())
async def get_questions(message: types.Message):
    kb = [
        [
            types.InlineKeyboardButton(text=id, callback_data=id)
        ] for id in range(len(await get_questions()))
    ]

    markup = types.InlineKeyboardMarkup(inline_keyboard=kb)

    await message.answer("Выберите вопрос:", reply_markup=markup)


@dp.message()
async def send_question(message: types.Message):
    await create_question(message.text, message.chat.id)
    await message.answer("Вопрос успешно отправлен!")


async def main():
    await create_tables()
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())