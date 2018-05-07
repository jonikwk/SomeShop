package bot

import (
	"database/sql"
	"fmt"
	"math/rand"

	cnf "../configuration"
	"../database"
	"github.com/fatih/color"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/glog"
	"googlemaps.github.io/maps"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (tgbot *TelegramBot) DeleteMessage(update tgbotapi.Update) {
	deleteMessage := tgbotapi.DeleteMessageConfig{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
	}
	tgbot.Token.Send(deleteMessage)
}

func (tgbot *TelegramBot) ChangeMessage(update tgbotapi.Update, db *sql.DB, messageID int, chatID int64, id int) {
	// id записи по имени из tables.catalog
	database.SetCurrentParnetId(db, chatID, id)  // в талице пользователей меняется id_parent
	markup := tgbot.SendSections(update, db, id) //тправка скций по
	edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
	tgbot.Token.Send(edit)
}

func (tgbot *TelegramBot) ChangeCurrentSection(update tgbotapi.Update, db *sql.DB, chatID int64) {
	idCurrent := database.GetCurrentParnetId(db, chatID)
	color.Yellow(fmt.Sprintln("ID CURRENT ЧТО ТОЛЬКО ЧТО СТАВИЛ: ", idCurrent))
	msg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
	msg.ReplyMarkup = tgbot.SendSections(update, db, idCurrent)
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

func GetMapsClient(config *cnf.Configuration) *maps.Client {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCPSKfYtsbLI1VmrfXYimmpZqDmIfZcEpQ" /*config.Settings.MapsApiKey*/))
	if err != nil {
		glog.Exit()
	}
	color.Red("HREEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE!!!")
	return c
}
