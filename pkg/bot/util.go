package bot

import (
	"database/sql"
	"fmt"

	"../database"
	"github.com/fatih/color"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tgbot *TelegramBot) DeleteMessage(update tgbotapi.Update) {
	deleteMessage := tgbotapi.DeleteMessageConfig{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
	}
	tgbot.Token.Send(deleteMessage)
}

func (tgbot *TelegramBot) ChangeMessage(update tgbotapi.Update, db *sql.DB, messageID int, chatID int64, id int) {
	database.SetCurrentParnetId(db, chatID, id)
	markup := tgbot.SendSections(update, db, id)
	edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
	tgbot.Token.Send(edit)
}

func (tgbot *TelegramBot) ChangeCurrentSection(update tgbotapi.Update, db *sql.DB, chatID int64) {
	idCurrent := database.GetCurrentParnetId(db, chatID)
	color.Yellow(fmt.Sprintln("ID CURRENT ЧТО ТОЛЬКО ЧТО СТАВИЛ: ", idCurrent))
	title := database.GetSectionTitle(db, idCurrent)
	msg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
	id := database.GetCatalogId(db, title)
	msg.ReplyMarkup = tgbot.SendSections(update, db, id)
	tgbot.Token.Send(msg)
}

func (tgbot *TelegramBot) IncreaseCurrentItem(db *sql.DB, chatID int64) {
	current := database.GetCurrentItem(db, chatID)
	current += 5
	database.SetCurrentItem(db, current, chatID)
}

func (tgbot *TelegramBot) DecreaseCurrentItem(db *sql.DB, chatID int64) {
	current := database.GetCurrentItem(db, chatID)
	current -= 5
	database.SetCurrentItem(db, current, chatID)
}
