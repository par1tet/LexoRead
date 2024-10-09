import os
import logging
import asyncio

from aiogram import Bot, Dispatcher, types
from aiogram.filters import Command, CommandStart, CommandObject
from dotenv import load_dotenv

load_dotenv()

TOKEN = os.getenv("TOKEN")

logging.basicConfig(level=logging.INFO)
bot = Bot(TOKEN)
dp = Dispatcher()

admins = [
    1340572920
]

questions = [
    {"question": "How i can create the book?", "chat_id": 1340572920}
]


@dp.callback_query()
async def btn_handler(call: types.CallbackQuery):
    data = questions[int(call.data)]
    question = data["question"]
    chat_id = f"t.me/testttstesbot?start={data["chat_id"]}"

    await call.message.answer(question + "\n\n" + chat_id)


@dp.message(CommandStart())
async def start(message: types.Message, command: CommandObject):
    if command.args and message.chat.id in admins:
        # TODO: соединить админа с вопросом
        await message.answer(command.args)
    else:
        await message.answer("Отправьте свой вопрос тех. поддержке")


@dp.message(Command("check"))
async def answering(message: types.Message):
    if message.chat.id in admins:
        kb = [[types.InlineKeyboardButton(text=str(id_), callback_data=str(id_))] for id_ in range(len(questions))]
        markup = types.InlineKeyboardMarkup(inline_keyboard=kb)
        await message.answer("Выберите вопрос:", reply_markup=markup)
    else:
        await message.answer("Нет прав.")


async def main():
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())