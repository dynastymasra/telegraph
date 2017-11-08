package telegraph

type (
	// InlineKeyboardMarkup This object represents an inline keyboard that appears right next to the message it belongs to.
	InlineKeyboardMarkup struct {
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	}

	// ReplyKeyboardMarkup This object represents a custom keyboard with reply options
	// (see Introduction to bots for details and examples).
	ReplyKeyboardMarkup struct {
		Keyboard        [][]KeyboardButton `json:"keyboard"`
		ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
		Selective       bool               `json:"selective,omitempty"`
	}

	// InlineKeyboardButton This object represents one button of an inline keyboard.
	// You must use exactly one of the optional fields.
	InlineKeyboardButton struct {
		Text                         string        `json:"text"`
		URL                          string        `json:"url,omitempty"`
		CallbackData                 string        `json:"callback_data,omitempty"`
		SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
		SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
		CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
		Pay                          bool          `json:"pay,omitempty"`
	}

	// KeyboardButton This object represents one button of the reply keyboard.
	// For simple text buttons String can be used instead of this object to specify text of the button.
	// Optional fields are mutually exclusive.
	KeyboardButton struct {
		Text            string `json:"text"`
		RequestContact  bool   `json:"request_contact,omitempty"`
		RequestLocation bool   `json:"request_location,omitempty"`
	}

	// ReplyKeyboardRemove Upon receiving a message with this object,
	// Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
	// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
	// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
	ReplyKeyboardRemove struct {
		RemoveKeyboard bool `json:"remove_keyboard"`
		Selective      bool `json:"selective,omitempty"`
	}
)
