package customkeyboard

import "encoding/json"

// WebAppInfo описывает информацию о Web App.
type WebAppInfo struct {
	URL string `json:"url"`
}

// InlineKeyboardButtonCustom — кнопка, поддерживающая WebApp.
type InlineKeyboardButtonCustom struct {
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// Можно добавить другие поля, например, callback_data, если понадобится.
}

// InlineKeyboardMarkupCustom описывает разметку клавиатуры.
type InlineKeyboardMarkupCustom struct {
	InlineKeyboard [][]InlineKeyboardButtonCustom `json:"inline_keyboard"`
}

// NewWebAppKeyboard создаёт клавиатуру с одной кнопкой Web App.
func NewWebAppKeyboard(buttonText, webAppURL string) InlineKeyboardMarkupCustom {
	button := InlineKeyboardButtonCustom{
		Text:   buttonText,
		WebApp: &WebAppInfo{URL: webAppURL},
	}
	keyboard := InlineKeyboardMarkupCustom{
		InlineKeyboard: [][]InlineKeyboardButtonCustom{
			{button},
		},
	}
	return keyboard
}

// MarshalKeyboard маршалит клавиатуру в JSON-строку.
func MarshalKeyboard(kb InlineKeyboardMarkupCustom) (string, error) {
	bytes, err := json.Marshal(kb)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
