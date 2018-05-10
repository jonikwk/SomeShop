package bot

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	cnf "../configuration"
	"../database"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/glog"
	"googlemaps.github.io/maps"
)

// AnalyzeUpdate -
func (tgbot *TelegramBot) AnalyzeUpdate(update tgbotapi.Update, db *sql.DB, config *cnf.Configuration) {
	switch {
	case update.CallbackQuery != nil:
		chatID := update.CallbackQuery.Message.Chat.ID
		messageID := update.CallbackQuery.Message.MessageID
		if database.IsUserInDatabase(chatID, db) == false {
			database.AddUser(db, chatID)
		}
		switch update.CallbackQuery.Data {
		case "XS":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "XS")
		case "S":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "S")
		case "M":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "M")
		case "L":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "L")
		case "XL":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "XL")
		case "XXL":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "XXL")
		case "XXXL":
			callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Вещь добавлена в корзину")
			tgbot.Token.AnswerCallbackQuery(callBack)
			tgbot.AddItemToOrder(update, db, "XXXL")
		case "Отзывы":
			tgbot.ShowReviews(update, db, chatID)
		case "Добавить отзыв":
			database.ActivateAddingReview(db, chatID)
			tgbot.AddReview(update, db, chatID)
		case "Одежда":
			id := database.GetCatalogID(db, "Одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Мужская одежда":
			id := database.GetCatalogID(db, "Мужская одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Женская одежда":
			id := database.GetCatalogID(db, "Женская одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Обувь":
			id := database.GetCatalogID(db, "Обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Женская обувь":
			id := database.GetCatalogID(db, "Женская обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Мужская обувь":
			id := database.GetCatalogID(db, "Мужская обувь")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Верхняя одежда":
			id := database.GetCatalogIDSameSections(db, chatID, "Верхняя одежда")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Футболки и майки":
			id := database.GetCatalogIDSameSections(db, chatID, "Футболки и майки")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Джемперы и толстовки":
			id := database.GetCatalogIDSameSections(db, chatID, "Джемперы и толстовки")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Блузки и рубашки":
			id := database.GetCatalogIDSameSections(db, chatID, "Блузки и рубашки")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Пиджаки и жакеты":
			id := database.GetCatalogIDSameSections(db, chatID, "Пиджаки и жакеты")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Брюки и джинсы":
			id := database.GetCatalogIDSameSections(db, chatID, "Брюки и джинсы")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Джемперы и свитеры":
			id := database.GetCatalogIDSameSections(db, chatID, "Джемперы и свитеры")
			tgbot.ChangeMessage(update, db, messageID, chatID, id)
		case "Блузки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Блузки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Рубашки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Рубашки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Брюки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Брюки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Джинсы":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Джинсы")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Футболки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Футболки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Платья":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Платья")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Юбки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Юбки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Жилеты":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Жилеты")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Комбинезоны":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Комбинезоны")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Куртки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Куртки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Майки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Майки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Пальто":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Пальто")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Шубы":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Шубы")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Дубленки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Дубленки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Бомберы":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Бомберы")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Парки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Парки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Плащи":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Плащи")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Джемперы":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Джемперы")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Свитеры":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Свитеры")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Пиджаки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Пиджаки")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Жакеты":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Жакеты")
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.SendItems(update, db, id)
		case "Толстовки":
			tgbot.DeleteMessage(update)
			id := database.GetCatalogIDSameSections(db, chatID, "Толстовки")
			database.SetCurrentParnetID(db, chatID, id)
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
			idCurrent := database.GetCurrentParnetID(db, chatID)
			id := database.GetParentID(db, idCurrent)
			database.SetCurrentParnetID(db, chatID, id)
			tgbot.ChangeCurrentSection(update, db, chatID)
		case "Назад куда то":
			tgbot.DeleteMessage(update)
			offset := database.GetCurrentItem(db, chatID)
			if offset > 0 {
				offset--
			}
			msg := tgbot.GetOrders(update, db, chatID, offset)
			database.SetCurrentItem(db, offset, chatID)
			tgbot.Token.Send(msg)
		case "Вперед куда то":
			tgbot.DeleteMessage(update)
			offset := database.GetCurrentItem(db, chatID)
			if offset < database.GetItemsInBucket(db, chatID)-1 {
				offset++
			}
			msg := tgbot.GetOrders(update, db, chatID, offset)
			database.SetCurrentItem(db, offset, chatID)
			tgbot.Token.Send(msg)
		case "Увеличить":
			tgbot.IncreaseItem(update, db, chatID)
			tgbot.DeleteMessage(update)
			offset := database.GetCurrentItem(db, chatID)
			msg := tgbot.GetOrders(update, db, chatID, offset)
			tgbot.Token.Send(msg)
		case "Уменьшить":
			tgbot.DecreaseItem(update, db, chatID)
			tgbot.DeleteMessage(update)
			offset := database.GetCurrentItem(db, chatID)
			msg := tgbot.GetOrders(update, db, chatID, offset)
			tgbot.Token.Send(msg)
		case "Удалить":
			tgbot.DeleteItem(update, db, chatID)
			tgbot.DeleteMessage(update)
			orderID := database.GetUserOrdersID(db, chatID)
			if database.GetUserOrdersCount(db, orderID) == 0 {
				msg := tgbotapi.NewMessage(chatID, "Ваша корзина пуста. Перейдите в каталог для выбора товаров")
				tgbot.Token.Send(msg)
			} else {
				offset := database.GetCurrentItem(db, chatID)
				msg := tgbot.GetOrders(update, db, chatID, offset)
				tgbot.Token.Send(msg)
			}
		case "Ещё":
			tgbot.DeleteMessage(update)
			idCurrent := database.GetCurrentParnetID(db, chatID)
			tgbot.IncreaseCurrentItem(db, chatID)
			tgbot.SendItems(update, db, idCurrent)
		}

	case update.Message != nil:
		chatID := update.Message.Chat.ID
		if database.IsUserInDatabase(chatID, db) == false {
			database.AddUser(db, chatID)
		}

		switch update.Message.Text {
		case "/start":
			tgbot.Greeting(update)
			tgbot.SendMenu(update)
		case "Корзина":
			orderID := database.GetUserOrdersID(db, chatID)
			if database.GetUserOrdersCount(db, orderID) == 0 {
				msg := tgbotapi.NewMessage(chatID, "Ваша корзина пуста. Перейдите в каталог для выбора товаров")
				tgbot.Token.Send(msg)
			} else {
				database.SetCurrentItem(db, 0, chatID)
				offset := database.GetCurrentItem(db, chatID)
				msg := tgbot.GetOrders(update, db, chatID, offset)
				tgbot.Token.Send(msg)
			}
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
		case "Отменить регистрацию":
			tgbot.CancelRegistration(update, db, chatID)
		case "Регистрация":
			if database.IsRegistrationCompleted(db, chatID) {
				msg := tgbotapi.NewMessage(chatID, "Вы уже зарегистрированы")
				tgbot.Token.Send(msg)
			} else {
				switch {
				case database.IsUserContainPhoneNumber(db, chatID) == false:
					tgbot.GetTelephoneNumber(update)
				case database.IsRegistrationCompleted(db, chatID) == false:
					tgbot.GetAddress(update, db)
				}
			}
		case "Да":
			database.CompleteRegistration(db, chatID)
			msg := tgbotapi.NewMessage(chatID, "Registration completed")
			tgbot.Token.Send(msg)
			tgbot.SendMenu(update)
		case "Новости":
			msg := tgbotapi.NewMessage(chatID, "На данный момент новостей нет")
			tgbot.Token.Send(msg)
		default:
			condition := database.IsUserContainPhoneNumber(db, chatID) == false && update.Message.Contact != nil
			switch {
			case condition == true:
				switch update.Message.Chat.ID != int64(update.Message.Contact.UserID) {
				case true:
					msg := tgbotapi.NewMessage(chatID,
						"Данный номер не является номером телефона, к которому привязан Ваш аккаунт. Нажмите подтвердить чтобы отправить свой номер телефона.")
					tgbot.Token.Send(msg)
				case false:
					phoneNumber := update.Message.Contact.PhoneNumber
					database.SetUserPhoneNumber(db, chatID, phoneNumber)
					tgbot.GetAddress(update, db)
				}
			case database.IsRegistrationCompleted(db, chatID) == false && database.IsUserContainPhoneNumber(db, chatID) == true:
				switch strings.Contains(strings.ToLower(update.Message.Text), "калуга") {
				case true:
					tgbot.IsAddresCorrect(update, db, config)
				case false:
					msg := tgbotapi.NewMessage(chatID, "К сожалению, это не очень похоже на адрес :( \nПовторите ввод снова.")
					tgbot.Token.Send(msg)
				}
			default:
				if database.GetAddingReview(db, chatID) {
					text := update.Message.Text
					tgbot.AddTextReview(db, chatID, text)
					database.DeactivateAddingReview(db, chatID)
					msg := tgbotapi.NewMessage(chatID, "review added")
					tgbot.Token.Send(msg)
				}
				/*if update.Message.Photo != nil {
					photo := *update.Message.Photo
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, photo[0].FileID)
					tgbot.Token.Send(msg)
					color.Red(photo[0].FileID)
				}*/
			}
		}
	}
}

//SendItems -
func (tgbot *TelegramBot) SendItems(update tgbotapi.Update, db *sql.DB, id int) {
	chatID := update.CallbackQuery.Message.Chat.ID
	offset := database.GetCurrentItem(db, chatID)
	items := database.GetItems(db, id, offset)
	sizesKeyboard := []tgbotapi.InlineKeyboardButton{}
	var photoMsg tgbotapi.PhotoConfig
	for _, item := range items {
		keyboard := tgbotapi.InlineKeyboardMarkup{}
		sizes := database.GetSizes(db, item.ID)
		photoMsg = tgbotapi.NewPhotoShare(chatID, item.Photo)
		photoMsg.Caption = fmt.Sprintf("%s\nЦена: %s\nЦвет: %s\n%s", item.Title, item.Price, item.Color, item.Description)
		if len(sizes) != 0 {
			for _, size := range sizes {
				button := tgbotapi.NewInlineKeyboardButtonData(size, size)
				sizesKeyboard = append(sizesKeyboard, button)
			}
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, sizesKeyboard)

			sizesKeyboard = []tgbotapi.InlineKeyboardButton{}
			review := tgbotapi.NewInlineKeyboardButtonData("Отзывы", "Отзывы")
			newReview := tgbotapi.NewInlineKeyboardButtonData("Добавить отзыв", "Добавить отзыв")
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{review})
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{newReview})
			photoMsg.ReplyMarkup = keyboard
			tgbot.Token.Send(photoMsg)
		} else {
			tgbot.Token.Send(photoMsg)
			msg := tgbotapi.NewMessage(chatID, "Данного товара нет в наличии")
			tgbot.Token.Send(msg)
		}
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

//SendSections -
func (tgbot *TelegramBot) SendSections(update tgbotapi.Update, db *sql.DB, id int) tgbotapi.InlineKeyboardMarkup {
	// id записи по имени из tables.catalog
	chatID := update.CallbackQuery.Message.Chat.ID
	offset := database.GetCurrentItem(db, chatID)    // возвращается число через сколько записей смотреть, offest
	recordsCount := database.GetRecordsCount(db, id) //количество записей в которй id_parent = id раздела
	sections := database.GetClothes(db, offset, id)  // возвращаются названия секций, у которых id_parent = id
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
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

//SendCatalog -
func (tgbot *TelegramBot) SendCatalog(update tgbotapi.Update, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	sections := database.GetRootSection(db)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, section := range sections {
		btn := tgbotapi.NewInlineKeyboardButtonData(section, section)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{btn})
	}
	return keyboard
}

//Greeting -
func (tgbot *TelegramBot) Greeting(update tgbotapi.Update) {
	firstName, chatID := update.Message.From.FirstName, update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Приветсвую Вас, %s", firstName))
	tgbot.Token.Send(msg)
}

//SendMenu -
func (tgbot *TelegramBot) SendMenu(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Главное меню:")
	catalog := tgbotapi.NewKeyboardButton("Каталог")
	bucket := tgbotapi.NewKeyboardButton("Корзина")
	registration := tgbotapi.NewKeyboardButton("Регистрация")
	news := tgbotapi.NewKeyboardButton("Новости")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{catalog, bucket}, {registration, news}}, ResizeKeyboard: true, OneTimeKeyboard: false}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

//SendMenuButton -
func (tgbot *TelegramBot) SendMenuButton(update tgbotapi.Update) tgbotapi.ReplyKeyboardMarkup {
	menu := tgbotapi.NewKeyboardButton("Главное меню")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{menu}}, ResizeKeyboard: true, OneTimeKeyboard: false}
	return keyboard
}

//GetTelephoneNumber -
func (tgbot *TelegramBot) GetTelephoneNumber(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintln(fmt.Sprintf("Телефон")))
	acceptButton, declineButton := tgbotapi.NewKeyboardButtonContact("Поделиться"), tgbotapi.NewKeyboardButton("Отменить регистрацию")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{acceptButton, declineButton}},
		ResizeKeyboard: true, OneTimeKeyboard: false}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

//GetAddress -
func (tgbot *TelegramBot) GetAddress(update tgbotapi.Update, db *sql.DB) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Адрес. Формат: \n Город, улица номер дома корпус/строение, квартира(если не частный дом) \nНапример: Калуга, Гагарина 13 б, 1\nКалуга, Гурьянова 59 корпус 3, 54")
	declineButton := tgbotapi.NewKeyboardButton("Отменить регистрацию")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{declineButton}}, ResizeKeyboard: true, OneTimeKeyboard: false}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

//IsAddresCorrect -
func (tgbot *TelegramBot) IsAddresCorrect(update tgbotapi.Update, db *sql.DB, config *cnf.Configuration) bool {
	client := GetMapsClient(config)
	var msg tgbotapi.MessageConfig
	chatID := update.Message.Chat.ID
	address := update.Message.Text
	r := &maps.GeocodingRequest{
		Address: address,
		Region:  "Россия",
	}
	resp, err := client.Geocode(context.Background(), r)
	if err != nil {
		glog.Exit()
	}

	if len(resp) == 0 {
		msg = tgbotapi.NewMessage(chatID, "К сожалению, данный адрес не найден. Проверьте правильность адреса и повторите ввод.")
		tgbot.Token.Send(msg)
		return false
	}

	status := resp[0].Geometry.LocationType
	switch status {
	case "RANGE_INTERPOLATED", "GEOMETRIC_CENTER", "APPROXIMATE":
		msg = tgbotapi.NewMessage(chatID, "К сожалению, я не смог точно определить ваш адрес. Проверьте правильность адреса и повторите ввод.")
		tgbot.Token.Send(msg)
		return false
	}
	tgbot.SendLocation(update, resp)
	database.AddAddress(db, chatID, update.Message.Text)
	tgbot.ConfirmAddress(update, db)
	return true
}

//SendLocation -
func (tgbot *TelegramBot) SendLocation(update tgbotapi.Update, resp []maps.GeocodingResult) {
	longtitude := resp[0].Geometry.Location.Lng
	lattitude := resp[0].Geometry.Location.Lat
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewLocation(chatID, lattitude, longtitude)
	tgbot.Token.Send(msg)
}

//ConfirmAddress -
func (tgbot *TelegramBot) ConfirmAddress(update tgbotapi.Update, db *sql.DB) {
	chatID := update.Message.Chat.ID
	address := database.GetAddress(db, chatID)
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Я нашел ваш дом. Нажмите Да, чтобы подтвердить адрес или введите новый, если вы указали неверный\n%s", address))
	acceptButton, declineButton := tgbotapi.NewKeyboardButton("Да"), tgbotapi.NewKeyboardButton("Отменить регистрацию")
	keyboard := tgbotapi.ReplyKeyboardMarkup{Keyboard: [][]tgbotapi.KeyboardButton{{acceptButton, declineButton}},
		ResizeKeyboard: true, OneTimeKeyboard: false}
	msg.ReplyMarkup = keyboard
	tgbot.Token.Send(msg)
}

// CancelRegistration -
func (tgbot *TelegramBot) CancelRegistration(update tgbotapi.Update, db *sql.DB, chatID int64) {
	database.SetUserInformationByDefault(db, chatID)
	msg := tgbotapi.NewMessage(chatID, "Регистрация отменена")
	tgbot.Token.Send(msg)
	tgbot.SendMenu(update)
}

//AddItemToOrder -
func (tgbot *TelegramBot) AddItemToOrder(update tgbotapi.Update, db *sql.DB, size string) {
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	chatID := update.CallbackQuery.Message.Chat.ID
	id := database.GetUserOrdersID(db, chatID)
	if id == 0 {
		number := RandStringBytes()
		database.AddOrder(db, number, chatID)
		id = database.GetUserOrdersID(db, chatID)
	}
	idSize := database.GetSizeID(db, size)
	database.AddItemToOrder(db, productID, id, idSize)
}

//GetOrders -
func (tgbot *TelegramBot) GetOrders(update tgbotapi.Update, db *sql.DB, chatID int64, offset int) tgbotapi.PhotoConfig {
	item := database.GetOrders(db, chatID, offset)
	msg := tgbotapi.NewPhotoShare(chatID, item.Photo)
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	delete := tgbotapi.NewInlineKeyboardButtonData("X", "Удалить")
	decrease := tgbotapi.NewInlineKeyboardButtonData("-", "Уменьшить")
	quantity := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d шт", item.Quantity), "default")
	increase := tgbotapi.NewInlineKeyboardButtonData("+", "Увеличить")
	back := tgbotapi.NewInlineKeyboardButtonData("<-", "Назад куда то")
	count := database.GetItemsInBucket(db, chatID)
	current := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d/%d", offset+1, count), "текущее количество")
	toward := tgbotapi.NewInlineKeyboardButtonData("->", "Вперед куда то")
	msg.Caption = fmt.Sprintf(" %s\nСтоимость: %d * %d = %d рублей\nЦвет: %s\nРазмер: %s",
		item.Title, item.Price, item.Quantity, item.Price*item.Quantity, item.Color, item.Size)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{delete, decrease, quantity, increase})
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{back, current, toward})
	msg.ReplyMarkup = keyboard
	return msg

}

//IncreaseItem -
func (tgbot *TelegramBot) IncreaseItem(update tgbotapi.Update, db *sql.DB, chatID int64) {
	caption := update.CallbackQuery.Message.Caption
	r, _ := regexp.Compile(`([0-9]{2}|[A-Z]{1,4})$`)
	size := r.FindString(caption)
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	orderID := database.GetUserOrdersID(db, chatID)
	sizeID := database.GetSizeID(db, size)
	database.AddItemToOrder(db, productID, orderID, sizeID)
	callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Количество товара увеличено")
	tgbot.Token.AnswerCallbackQuery(callBack)
}

//DeleteItem -
func (tgbot *TelegramBot) DeleteItem(update tgbotapi.Update, db *sql.DB, chatID int64) {
	caption := update.CallbackQuery.Message.Caption
	r, _ := regexp.Compile(`([0-9]{2}|[A-Z]{1,4})$`)
	size := r.FindString(caption)
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	orderID := database.GetUserOrdersID(db, chatID)
	sizeID := database.GetSizeID(db, size)
	database.DeleteItemFromOrder(db, productID, orderID, sizeID)
	offset := database.GetCurrentItem(db, chatID)
	if offset > 0 {
		offset--
	}
	database.SetCurrentItem(db, offset, chatID)
	callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Товар удален")
	tgbot.Token.AnswerCallbackQuery(callBack)
}

//DecreaseItem -
func (tgbot *TelegramBot) DecreaseItem(update tgbotapi.Update, db *sql.DB, chatID int64) {
	caption := update.CallbackQuery.Message.Caption
	r, _ := regexp.Compile(`([0-9]{2}|[A-Z]{1,4})$`)
	size := r.FindString(caption)
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	orderID := database.GetUserOrdersID(db, chatID)
	sizeID := database.GetSizeID(db, size)
	database.ChangeQuantityItemToOrder(db, productID, orderID, sizeID, -1)
	callBack := tgbotapi.NewCallback(update.CallbackQuery.ID, "Количество товара уменьшено")
	tgbot.Token.AnswerCallbackQuery(callBack)
}

//AddReview -
func (tgbot *TelegramBot) AddReview(update tgbotapi.Update, db *sql.DB, chatID int64) {
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	var name string
	if update.CallbackQuery.From.UserName != "" {
		name = update.CallbackQuery.From.UserName
	} else {
		name = update.CallbackQuery.From.FirstName + " " + update.CallbackQuery.From.LastName
	}
	database.AddAuthorReview(db, chatID, productID, name)
	msg := tgbotapi.NewMessage(chatID, "Your review:")
	tgbot.Token.Send(msg)
}

//AddTextReview -
func (tgbot *TelegramBot) AddTextReview(db *sql.DB, chatID int64, text string) {
	database.AddTextReview(db, chatID, text)
}

//ShowReviews -
func (tgbot *TelegramBot) ShowReviews(update tgbotapi.Update, db *sql.DB, chatID int64) {
	photo := *update.CallbackQuery.Message.Photo
	photoID := photo[0].FileID
	productID := database.GetProductID(db, photoID)
	items := database.GetReviews(db, productID)
	if len(items) != 0 {
		for _, item := range items {
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s\n Date: %s\n %s", item.Name, item.Date, item.Description))
			tgbot.Token.Send(msg)
		}
	} else {
		msg := tgbotapi.NewMessage(chatID, "Threre are no reviews on this item.")
		tgbot.Token.Send(msg)
	}
}
