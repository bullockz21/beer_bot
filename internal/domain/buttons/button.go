package buttons

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Button структура для кнопки
type Button struct {
	Text string
	Data string
}

// InlineKeyboard возвращает готовую инлайн-клавиатуру
func InlineKeyboard(buttons ...Button) tgbotapi.InlineKeyboardMarkup {
	rows := make([]tgbotapi.InlineKeyboardButton, len(buttons))
	for i, btn := range buttons {
		rows[i] = tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows)
}

func InlineKeyboardColumn(buttons ...Button) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, btn := range buttons {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data),
		)
		rows = append(rows, row)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

// Предопределенные кнопки
var (
	MenuButton       = Button{Text: "📜 Меню", Data: "menu"}
	PromotionsButton = Button{Text: "🔥 Акции", Data: "promotions"}
	ReviewsButton    = Button{Text: "⭐ Отзывы", Data: "reviews"}
)

var (
	ShawarmaButton = Button{Text: "Шаурма", Data: "shawarma"}
	DrinksButton   = Button{Text: "Напитки", Data: "drinks"}
	DessertsButton = Button{Text: "Десерты", Data: "desserts"}
	BackButton     = Button{Text: "Назад", Data: "back"}
)
