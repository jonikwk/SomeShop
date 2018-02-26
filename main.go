package main

import (
	"./pkg/bot"
	cnf "./pkg/configuration"
	"./pkg/database"
)

func main() {
	var config = new(cnf.Configuration)
	config.ParseConfigurationFile()

	psqlInfo := database.GetConnectionString(config)
	db := database.OpenDB(config, psqlInfo)
	defer db.Close()
	db.Ping()

	var bot = new(bot.TelegramBot)
	bot.Init(config)
	bot.Start(db, config)
}
