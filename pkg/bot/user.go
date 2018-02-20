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
			id := database.GetCatalogId(db, "Одежда") //возвращается id записи по имени
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		caеуse "Мужская одежда":
			id := database.GetCatalogId(db, "Мужская одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Женская одежда":
			id := database.GetCatalogId(db, "Женская одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Обувь":
			id := database.GetCatalogId(db, "Обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Женская обувь":
			id := database.GetCatalogId(db, "Женская обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Мужская обувь":
			id := database.GetCatalogId(db, "Мужская обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Верхняя одежда":
			id := database.GetCatalogIDSameSections(db, chatID, "Верхняя одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Футболки и майки":
			id := database.GetCatalogIDSameSections(db, chatID, "Футболки и майки")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Футболки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Футболки")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)
		case "Платья":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Платья")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)
		case "Юбки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Юбки")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)
		case "Жилеты":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Жилеты")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)
		case "Комбинезоны":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Комбинезоны")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)

		case "Майки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Майки")
			database.SetCurrentParnetId(db, chatID, id) // в талице пользователей меняется id_parent
			tgbot.SendItems(update, db, id)
		case "Каталог вперед":
			tgbot.DeleteMessage(update)
			tgbot.IncreaseCurrentItem(db, chatID)
			tgbot.ChangeCurrentSection(update, db, chatID)
		case "Каталог назад":
			tgbot.DeleteMessage(update)
			tgbot.DecreaseCurrentItem(db, chatID)
			tgbot.ChangeCurrentSection(update, db, chatID)
		case "Назад":
			tgbot.DeleteMessage(update)
			idCurrent := database.GetCurrentParnetId(db, chatID)
			color.Green(fmt.Sprintln("ID CURRENT: ", idCurrent))
			id := database.GetParentID(db, idCurrent)
			color.Green(fmt.Sprintln("ID PARENT: ", id))
			database.SetCurrentParnetId(db, chatID, id)
			tgbot.ChangeCurrentSection(update, db, chatID)
		case "Ещё":
			tgbot.DeleteMessage(update)
			idCurrent := database.GetCurrentParnetId(db, chatID)
			color.Green(fmt.Sprintln("ID CURRENT: ", idCurrent))
			tgbot.IncreaseCurrentItem(db, chatID)
			tgbot.SendItems(update, db, idCurrent)
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
			menuMsg := tgbotapi.NewMessage(chatID, "<i>Каталог:</i>")
			menuMsg.ParseMode = "HTML"
			menuMsg.ReplyMarkup = tgbot.SendMenuButton(update)
			catalogMsg := tgbotapi.NewMessage(chatID, "Выберите раздел:")
			catalogMsg.ReplyMarkup = tgbot.SendCatalog(update, db)
			tgbot.Token.Send(menuMsg)
			tgbot.Token.Send(catalogMsg)
		case "Главное меню":
			tgbot.SendMenu(update)
		case "Регистрация":
			msg := tgbotapi.NewPhotoShare(chatID, "AgADAgAD66gxG5FEUUhyy2GRiLwx8s8MnA4ABCetSue57gYe7JABAAEC")
			msg.Caption = "2345678"
			tgbot.Token.Send(msg)
		default:
			if update.Message.Photo != nil {
				photo := *update.Message.Photo
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, photo[0].FileID)
				tgbot.Token.Send(msg)
				color.Red(photo[0].FileID)
			}
		}
	}
}

func (tgbot *TelegramBot) SendItems(update tgbotapi.Update, db *sql.DB, id int) {
	color.Red("HERE!!!!!!!")
	chatID := update.CallbackQuery.Message.Chat.ID
	offset := database.GetCurrentItem(db, chatID)
	color.Yellow(fmt.Sprintln("OFFSET: ", offset))
	items := database.GetItems(db, id, offset)
	color.Green(fmt.Sprintln("ITEMS: ", items))
	for _, item := range items {
		keyboard := tgbotapi.InlineKeyboardMarkup{}
		bucket := tgbotapi.NewInlineKeyboardButtonData("В корзину", "В корзину")
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{bucket})
		msg := tgbotapi.NewPhotoShare(chatID, item.Photo)
		msg.Caption = fmt.Sprintf("%s\nЦена: %s\nЦвет: %s\n%s", item.Title, item.Price, item.Color, item.Description)
		msg.ReplyMarkup = keyboard
		tgbot.Token.Send(msg)
	}
	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	count := database.GetItemsCount(db, id)
	another := tgbotapi.NewInlineKeyboardButtonData("Ещё", "Ещё")
	back := tgbotapi.NewInlineKeyboardButtonData("К каталогу", "Назад")
	if offset+5 >= count {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back})
	} else {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{another, back})
	}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

func (tgbot *TelegramBot) SendSections(update tgbotapi.Update, db *sql.DB, id int) tgbotapi.InlineKeyboardMarkup {
	// id записи по имени из tables.catalog
	chatID := update.CallbackQuery.Message.Chat.ID
	offset := database.GetCurrentItem(db, chatID)    // возвращается число через сколько записей смотреть, offest
	recordsCount := database.GetRecordsCount(db, id) //количество записей в которй id_parent = id раздела
	sections := database.GetClothes(db, offset, id)  // возвращаются названия секций, у которых id_parent = id
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		color.Red(section)
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	if id == 1 || id == 2 {
		back := tgbotapi.NewInlineKeyboardButtonData("🔼", "Назад")
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back})

	} else if id > 2 {
		back := tgbotapi.NewInlineKeyboardButtonData("🔼", "Назад")
		right := tgbotapi.NewInlineKeyboardButtonData("➡️", "Каталог вперед")
		left := tgbotapi.NewInlineKeyboardButtonData("⬅️", "Каталог назад")
		switch {
		case recordsCount <= 5:
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back})
		case recordsCount-offset <= 5:
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{left, back})
		case offset == 0:
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back, right})
		case offset > 0:
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{left, back, right})

		}
	}

	return keyboard
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

func (tgbot *TelegramBot) SendMenuButton(update tgbotapi.Update) tgbotapi.ReplyKeyboardMarkup {
	menu := tgbotapi.NewKeyboardButton("Главное меню")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{menu}}, ResizeKeyboard: true, OneTimeKeyboard: true}
	return keyboard
}
