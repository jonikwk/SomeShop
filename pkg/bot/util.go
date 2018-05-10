package bot

import (
	"database/sql"
	"math/rand"

	cnf "../configuration"
	"../database"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/glog"
	"googlemaps.github.io/maps"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

//RandStringBytes -
func RandStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//DeleteMessage -
func (tgbot *TelegramBot) DeleteMessage(update tgbotapi.Update) {
	deleteMessage := tgbotapi.DeleteMessageConfig{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
	}
	tgbot.Token.Send(deleteMessage)
}

//ChangeMessage -
func (tgbot *TelegramBot) ChangeMessage(update tgbotapi.Update, db *sql.DB, messageID int, chatID int64, id int) {
	database.SetCurrentParnetID(db, chatID, id)
	markup := tgbot.SendSections(update, db, id)
	edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
	tgbot.Token.Send(edit)
}

//ChangeCurrentSection -
func (tgbot *TelegramBot) ChangeCurrentSection(update tgbotapi.Update, db *sql.DB, chatID int64) {
	idCurrent := database.GetCurrentParnetID(db, chatID)
	msg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
	msg.ReplyMarkup = tgbot.SendSections(update, db, idCurrent)
	tgbot.Token.Send(msg)
}

//IncreaseCurrentItem -
func (tgbot *TelegramBot) IncreaseCurrentItem(db *sql.DB, chatID int64) {
	current := database.GetCurrentItem(db, chatID)
	current += 5
	database.SetCurrentItem(db, current, chatID)
}

//DecreaseCurrentItem -
func (tgbot *TelegramBot) DecreaseCurrentItem(db *sql.DB, chatID int64) {
	current := database.GetCurrentItem(db, chatID)
	current -= 5
	database.SetCurrentItem(db, current, chatID)
}

//GetMapsClient -
func GetMapsClient(config *cnf.Configuration) *maps.Client {
	c, err := maps.NewClient(maps.WithAPIKey(config.Settings.MapsAPIKey))
	if err != nil {
		glog.Exit()
	}
	return c
}
