package bot

import (
	"database/sql"
	"fmt"

	"../database"
	"github.com/fatih/color"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tgbot *TelegramBot) AnalyzeUpdate(update tgbotapi.Update, db *sql.DB) {
	switch {
	case update.CallbackQuery != nil:
		chatID := update.CallbackQuery.Message.Chat.ID
		messageID := update.CallbackQuery.Message.MessageID
		if database.IsUserInDatabase(chatID, db) == false {
			color.Red(fmt.Sprintln("CallBACL: ", chatID))
			database.AddUser(db, chatID)
		} //ОБНУЛИТЬ ЗНАЧЕНИЯ

		switch update.CallbackQuery.Data {
		case "Одежда":
			markup := tgbot.SendClothes(update, db)
			edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
			tgbot.Token.Send(edit)
		case "К каталогу":
			markup := tgbot.SendCatalog(update, db)
			edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
			tgbot.Token.Send(edit)
		case "Женская одежда":
			id := database.GetCatalogId(db, "Женская одежда")
			database.SetCurrentParnetId(db, chatID, id)
			database.SetCurrentItemByDefault(db, chatID)

			markup := tgbot.SendSectionItems(update, db, id)
			edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
			tgbot.Token.Send(edit)
		case "Мужская-женская":
			markup := tgbot.SendClothes(update, db)
			edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
			tgbot.Token.Send(edit)
		case "Мужская одежда":
			id := database.GetCatalogId(db, "Мужская одежда")
			color.Red(fmt.Sprintln(id))
			database.SetCurrentParnetId(db, chatID, id)
			database.SetCurrentItemByDefault(db, chatID)

			markup := tgbot.SendSectionItems(update, db, id)
			edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
			tgbot.Token.Send(edit)
		case "Каталог вперед":
			deleteMessage := tgbotapi.DeleteMessageConfig{}
			deleteMessage.ChatID = chatID
			deleteMessage.MessageID = messageID
			tgbot.Token.Send(deleteMessage)
			current := database.GetCurrentItem(db, chatID)
			current += 5
			database.SetCurrentItem(db, current, chatID)

			idCurrent := database.GetCurrentParnetId(db, chatID)
			title := database.GetSectionTitle(db, idCurrent)
			msg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
			id := database.GetCatalogId(db, title)
			msg.ReplyMarkup = tgbot.SendSectionItems(update, db, id)
			tgbot.Token.Send(msg)
		case "Каталог назад":
			deleteMessage := tgbotapi.DeleteMessageConfig{}
			deleteMessage.ChatID = chatID
			deleteMessage.MessageID = messageID
			tgbot.Token.Send(deleteMessage)
			current := database.GetCurrentItem(db, chatID)
			current -= 5
			database.SetCurrentItem(db, current, chatID)

			idCurrent := database.GetCurrentParnetId(db, chatID)
			title := database.GetSectionTitle(db, idCurrent)
			msg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
			id := database.GetCatalogId(db, title)
			msg.ReplyMarkup = tgbot.SendSectionItems(update, db, id)
			tgbot.Token.Send(msg)
		}

	case update.Message != nil:
		chatID := update.Message.Chat.ID
		if database.IsUserInDatabase(chatID, db) == false {
			color.Red(fmt.Sprintln("USUAL: ", chatID))
			database.AddUser(db, chatID)
		}

		switch update.Message.Text {
		case "/start":
			tgbot.Greeting(update)
			tgbot.SendMenu(update)
		case "Каталог":
			menuMsg := tgbotapi.NewMessage(chatID, "Каталог:")
			menuMsg.ReplyMarkup = tgbot.SendMenuButton(update)
			catalogMsg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
			catalogMsg.ReplyMarkup = tgbot.SendCatalog(update, db)
			tgbot.Token.Send(menuMsg)
			tgbot.Token.Send(catalogMsg)
		case "Главное меню":
			tgbot.SendMenu(update)
		default:
			tgbot.SendMenu(update)
		}
	}
}

func (tgbot *TelegramBot) Greeting(update tgbotapi.Update) {
	firstName, chatID := update.Message.From.FirstName, update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Приветсвую Вас, %s", firstName))
	tgbot.Token.Send(msg)
}

func (tgbot *TelegramBot) SendMenu(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Главное меню:")
	catalog := tgbotapi.NewKeyboardButton("Каталог")
	bucket := tgbotapi.NewKeyboardButton("Корзина")
	registration := tgbotapi.NewKeyboardButton("Регистрация")
	news := tgbotapi.NewKeyboardButton("Новости")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{catalog, bucket}, {registration, news}}, ResizeKeyboard: true, OneTimeKeyboard: true}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

func (tgbot *TelegramBot) SendCatalog(update tgbotapi.Update, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	sections := database.GetRootSection(db)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	return keyboard
}

func (tgbot *TelegramBot) SendClothes(update tgbotapi.Update, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	sections := database.GetClothesSection(db)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	back := tgbotapi.NewInlineKeyboardButtonData("Влдислав, назад", "К каталогу")
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back})
	return keyboard
}

func (tgbot *TelegramBot) SendMenuButton(update tgbotapi.Update) tgbotapi.ReplyKeyboardMarkup {
	menu := tgbotapi.NewKeyboardButton("Главное меню")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{menu}}, ResizeKeyboard: true, OneTimeKeyboard: true}
	return keyboard
}

func (tgbot *TelegramBot) SendManClothes(update tgbotapi.Update, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	sections := database.GetManClothes(db)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	back := tgbotapi.NewInlineKeyboardButtonData("Назад", "Мужская-женская")
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back})
	return keyboard
}

//общая функция отправки одежды
func (tgbot *TelegramBot) SendSectionItems(update tgbotapi.Update, db *sql.DB, id int) tgbotapi.InlineKeyboardMarkup {
	chatID := update.CallbackQuery.Message.Chat.ID
	current := database.GetCurrentItem(db, chatID)
	recordsCount := database.GetRecordsCount(db, id)
	color.Green(fmt.Sprintln("CURENT: ", current))
	sections := database.GetClothes(db, current, id)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	back := tgbotapi.NewInlineKeyboardButtonData("🔼", "Мужская-женская")    //"Мужская-женская"
	forward := tgbotapi.NewInlineKeyboardButtonData("➡️", "Каталог вперед") //каталог одежды назад вперед
	torward := tgbotapi.NewInlineKeyboardButtonData("⬅️", "Каталог назад")  //каталог одежды назад вперед
	switch {
	case recordsCount-current <= 5:
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{torward, back})
	case current == 0:
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back, forward})
	case current > 0:
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{torward, back, forward})
	}
	return keyboard
}
