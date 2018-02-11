package bot

import (
	"database/sql"

	cnf "../configuration"
	"github.com/fatih/color"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/glog"
)

//TelegramBot -> Представляет общую структуру бота
type TelegramBot struct {
	Token   *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

//Init -> Начальная инициализация бота необходимыми параметрами
func (tgbot *TelegramBot) Init(config *cnf.Configuration) {
	botAPI, err := tgbotapi.NewBotAPI("409034602:AAE7Tbs2-7DF6h5h8RJDn47KFzPc9lGs5ic")
	if err != nil {
		glog.Exit(err)
	}
	botAPI.Debug = true
	color.Green("Autorized on account %s", botAPI.Self.UserName)

	tgbot.Token = botAPI
	botUpdate := tgbotapi.NewUpdate(config.Settings.UpdateOfSet)
	botUpdate.Timeout = config.Settings.UpdateTimeout
	botUpdates, err := tgbot.Token.GetUpdatesChan(botUpdate)
	if err != nil {
		glog.Exit(err)
	}
	tgbot.Updates = botUpdates
}

//Start -> Запуск цикла поиска обновлений
func (tgbot *TelegramBot) Start(db *sql.DB) {
	for update := range tgbot.Updates {
		tgbot.AnalyzeUpdate(update, db)
	}
}
