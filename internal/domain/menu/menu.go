package menu

// Доменное представление кнопки
type Button struct {
	Text    string
	Command string
}

// Доменное меню
type Menu struct {
	Title   string
	Buttons [][]Button
}

func StartMenu() *Menu {
	return &Menu{
		Title: "Главное меню",
		Buttons: [][]Button{
			{
				{Text: "📝 Отзывы", Command: "reviews"},
				{Text: "🎁 Промокоды", Command: "promocodes"},
			},
		},
	}
}
