package telegraph

type (

	// InlineQuery This object represents an incoming inline query. When the user sends an empty query,
	// your bot could return some default or trending results.
	InlineQuery struct {
		ID       string    `json:"id"`
		From     User      `json:"from"`
		Location *Location `json:"location,omitempty"`
		Query    string    `json:"query"`
		Offset   string    `json:"offset"`
	}

	// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
	ChosenInlineResult struct {
		ResultID        string    `json:"result_id"`
		From            User      `json:"from"`
		Location        *Location `json:"location,omitempty"`
		InlineMessageID string    `json:"inline_message_id,omitempty"`
		Query           string    `json:"query"`
	}

	// CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard.
	// If the button that originated the query was attached to a message sent by the bot,
	// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
	// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
	CallbackQuery struct {
		ID              string   `json:"id"`
		From            User     `json:"from"`
		Message         *Message `json:"message,omitempty"`
		InlineMessageID string   `json:"inline_message_id,omitempty"`
		ChatInstance    string   `json:"chat_instance"`
		Data            string   `json:"data,omitempty"`
		GameShortName   string   `json:"game_short_name,omitempty"`
	}

	// ShippingQuery This object contains information about an incoming shipping query.
	ShippingQuery struct {
		ID              string          `json:"id"`
		From            User            `json:"from"`
		InvoicePayload  string          `json:"invoice_payload"`
		ShippingAddress ShippingAddress `json:"shipping_address"`
	}

	// PreCheckoutQuery This object contains information about an incoming pre-checkout query.
	PreCheckoutQuery struct {
		ID               string `json:"id"`
		From             User   `json:"from"`
		Currency         string `json:"currency"`
		TotalAmount      int64  `json:"total_amount"`
		InvoicePayload   string `json:"invoice_payload"`
		ShippingOptionID string `json:"shipping_option_id,omitempty"`
		OrderInfo        string `json:"order_info,omitempty"`
	}

	// ForceReply Upon receiving a message with this object,
	// Telegram clients will display a reply interface to the user
	// (act as if the user has selected the bot‘s message and tapped ’Reply').
	// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
	ForceReply struct {
		ForceReply bool `json:"force_reply"`
		Selective  bool `json:"selective,omitempty"`
	}

	// ReplyMarkup Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup struct {
		*InlineKeyboardMarkup
		*ReplyKeyboardMarkup
		*ReplyKeyboardRemove
		*ForceReply
	}

	// PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		FileSize int    `json:"file_size"`
	}
)
