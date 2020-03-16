package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type OrderResp struct {
	Product string
	Amount  int
}

type Order struct {
	State      int // 0-name, 1-tel,  2-prod, 3-amount
	Name       string
	Telephone  string
	Product    string
	DateCreate string
	Amount     int
	CustomerID int
	ProductID  int
}
type Customer struct {
	ID           int    `db:"id"`
	CustomerName string `db:"customer_name"`
	Phone        string `db:"phone"`
	ChatID       int    `db:"chat_id"`
}

var productMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🥖 батон"),
		tgbotapi.NewKeyboardButton("🍞 кирпич"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🥐 булка"),
		tgbotapi.NewKeyboardButton("🍞 ржаной"),
	),
)
