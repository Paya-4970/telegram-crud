from telegram import Update, InlineKeyboardMarkup
from telegram.ext import ApplicationBuilder, CommandHandler, filters, ContextTypes, MessageHandler, ConversationHandler
from db import insert_food
from loadenv import TOKEN_BOT

NAME, INFO, PRICE = range(3)

# //////////// STRAT PART /////////////

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="Hello, I'm Telegram Crud Bot!",
        reply_to_message_id=update.message.message_id,
    )

async def echo(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    text = context.args[0]
    if context.args[0]:
        text = " ".join(context.args)
    elif context.args[0] == " " or context.args[0] == None:
        text = "You didn’t provide any text!"
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text=text,
        reply_to_message_id=update.message.message_id,
    )

async def help_command(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    help_text = (
        "🤖 *Telegram CRUD Bot Help*\n\n"
        "/start - Start the bot\n"
        "/food - Add a new food (name, info, price)\n"
        "/echo <text> - Echo the text back\n"
        "/help - Show this help message\n\n"
        "Workflow for /food:\n"
        "1️⃣ Send /food\n"
        "2️⃣ Enter the food name\n"
        "3️⃣ Enter the info/description\n"
        "4️⃣ Enter the price\n"
        "The bot will save it to the database and confirm."
    )
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text=help_text,
        reply_to_message_id=update.message.message_id,
    )


# ///////////////// FOOD PART //////////////////
async def get_food(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="Ok Now Send Me The Name Of The Food You Want!",
        reply_to_message_id=update.message.message_id,
    )
    return NAME

async def get_name_food(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int :
    context.user_data["name"] = update.message.text
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text= "Alright Your Name Saved Now Write The Info!",
        reply_to_message_id=update.message.message_id,
    )
    return INFO

async def get_info_food(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    context.user_data["info"] = update.message.text
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="Alright Your Info Saved Now Write The Price!",
        reply_to_message_id=update.message.message_id,
    )
    return PRICE

async def get_price_food(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    price_text = update.message.text
    try:
        price = int(price_text)
    except ValueError:
        await update.message.reply_text("Price must be a number! Start again with /food")
        return ConversationHandler.END

    name = context.user_data.get("name")
    info = context.user_data.get("info")

    try:
        food_id = insert_food(name, info, price)
        await update.message.reply_text(f"Food saved successfully! ID: {food_id}")
    except Exception as e:
        await update.message.reply_text(f"Database error: {e}")

    context.user_data.clear()

    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="Alright Your Price Saved\n Thanks!",
        reply_to_message_id=update.message.message_id,
    )
    return ConversationHandler.END

async def cancel_handler(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="Ok By\n",
        reply_to_message_id=update.message.message_id,
    )
    return ConversationHandler.END


# /////////////// MAIN PART /////////////////

if __name__ == '__main__':
    bot = ApplicationBuilder().token(TOKEN_BOT).build()

    conv = ConversationHandler(
        entry_points=[CommandHandler("food",get_food)],
        states={
            NAME : [MessageHandler(filters.TEXT & ~filters.COMMAND, get_name_food)],
            INFO : [MessageHandler(filters.TEXT & ~filters.COMMAND, get_info_food)],
            PRICE : [MessageHandler(filters.TEXT & ~filters.COMMAND, get_price_food)]
        },
        fallbacks = [MessageHandler(filters.ALL, cancel_handler)],
        allow_reentry = True 
    )
    bot.add_handler(conv)

    help_handler = CommandHandler("help", help_command)
    bot.add_handler(help_handler)
    start_handelr = CommandHandler("start", start)
    bot.add_handler(start_handelr)
    echo_handler = CommandHandler("echo", echo)
    bot.add_handler(echo_handler)

    bot.run_polling()
