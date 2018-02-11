package bot

import (
	"database/sql"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tgbot *TelegramBot) AnalyzeUpdate(update tgbotapi.Update, db *sql.DB) {
	switch {
	case update.Message != nil:
	}
}
