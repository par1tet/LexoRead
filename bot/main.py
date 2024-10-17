import logging
import asyncio

from aiogram import Bot, Dispatcher, types
from aiogram.filters import Command, CommandStart, CommandObject, BaseFilter

from config import TOKEN
from database.core import (
    create_tables,
    create_question,
    get_questions,
    leave_from_question,
    get_question,
    get_admins,
    get_chat_id_questioner
)


class CommandNonAdmin(BaseFilter):
    async def __call__(self, message: types.Message):
        if any((i in message.text for i in ("/check", "/leave", "/cancel", "/admin"))):
            return True 


class AdminFilter(BaseFilter):
    async def __call__(self, message: types.CallbackQuery | types.Message) -> bool:
        admins = await get_admins()
        if isinstance(message, types.Message):
            return message.chat.id in admins 
        return message.message.chat.id in admins


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


@dp.callback_query(AdminFilter())
async def btn_handler(call: types.CallbackQuery):
    if "." in call.data:
        page = str(int(call.data.split(".")[0]) + ((1 if "+" in call.data else -1) if ".0" not in call.data else 0))
        if int(page) < 0:
            page = 0
        kb = [
            *[[
                types.InlineKeyboardButton(text=id, callback_data=f"{id}|{page}")
            ] for id in await get_questions(int(page))],
            [
                types.InlineKeyboardButton(text="<", callback_data=f"{page}.-1"),
                types.InlineKeyboardButton(text=str(int(page) + 1), callback_data="-"),
                types.InlineKeyboardButton(text=">", callback_data=f"{page}.+1")
            ]
        ]
        markup = types.InlineKeyboardMarkup(inline_keyboard=kb)
        await call.message.edit_reply_markup(reply_markup=markup)
        return


    if call.data == "-":
        await call.answer("И чего ты этим добиваешься?")
        return

    chat_id, page = call.data.split("|")

    question = await get_question(chat_id)
    text = "Нажмите чтобы ответить:"

    kb = [[
        types.InlineKeyboardButton(
            text="Ответить", url=f"t.me/{(await bot.get_me()).username}?start={question.id}")
        ],[
        types.InlineKeyboardButton(
            text="Назад", callback_data=f"{page}.0")
    ]]

    markup = types.InlineKeyboardMarkup(inline_keyboard=kb)

    await call.message.edit_text(text=(question.text + "\n\n" + text), reply_markup=markup)


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
async def check_questions(message: types.Message):
    kb = [
        *([
            types.InlineKeyboardButton(text=id, callback_data=f"{id}|{0}")
        ] for id in await get_questions(0)),
        [
            types.InlineKeyboardButton(text="<", callback_data="0.-1"),
            types.InlineKeyboardButton(text="1", callback_data="-"),
            types.InlineKeyboardButton(text=">", callback_data="0.+1")
        ]
    ]

    markup = types.InlineKeyboardMarkup(inline_keyboard=kb)
    await message.answer("Выберите вопрос:", reply_markup=markup)

@dp.message(CommandNonAdmin())
async def nonadmins(message: types.Message):
    await message.answer("Нет прав.")


@dp.message(AdminFilter(), ConnectionFilter())
async def send_to_question(message: types.Message):
    chat_id = await get_chat_id_questioner(message.chat.id)
    await bot.send_message(chat_id, message.text)


@dp.message()
async def send_question(message: types.Message):
    if message.text or message.caption:
        await create_question(message.text, message.chat.id)
        await message.answer("Вопрос успешно отправлен!")
        print(f"НОВЫЙ ВОПРОС ОТ {message.from_user.username}: {message.text}")
        return
    await message.answer("Отправьте текстовое сообщение.")
    print(f"{message.from_user.username} попытался отправить стикер, гифку, или медиа контент")


async def main():
    await create_tables()
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
