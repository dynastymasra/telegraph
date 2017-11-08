package telegraph

type (
	ChatType  string
	ParseMode string
)

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML               = "HTML"
)

type (
	SendMessage struct {
		ChatID                string       `json:"chat_id"`
		Text                  string       `json:"text,omitempty"`
		ParseMode             ParseMode    `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
		DisableNotification   bool         `json:"disable_notification,omitempty"`
		ReplyMessageID        int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint              string       `json:"-"`
	}
)

// NewTextMessage Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// Text of the message to be sent
func NewTextMessage(chatID, text string) *SendMessage {
	return &SendMessage{
		ChatID:   chatID,
		Text:     text,
		endpoint: EndpointSendMessage,
	}
}

// SetParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic,
// fixed-width text or inline URLs in your bot's message.
func (message *SendMessage) SetParseMode(mode ParseMode) *SendMessage {
	message.ParseMode = mode
	return message
}

// SetDisableWebPagePreview Disables link previews for links in this message
func (message *SendMessage) SetDisableWebPagePreview(disable bool) *SendMessage {
	message.DisableWebPagePreview = disable
	return message
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *SendMessage) SetDisableNotification(disable bool) *SendMessage {
	message.DisableNotification = disable
	return message
}

// ReplyMessageToID If the message is a reply, ID of the original message
func (message *SendMessage) SetReplyMessageToId(id int64) *SendMessage {
	message.ReplyMessageID = id
	return message
}

// SetForceReply
func (message *SendMessage) SetForceReply(reply ForceReply) *SendMessage {
	message.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return message
}

// SetInlineKeyboardMarkup
func (message *SendMessage) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendMessage {
	message.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return message
}

// SetReplyKeyboardMarkup
func (message *SendMessage) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendMessage {
	message.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return message
}

// SetReplyKeyboardRemove
func (message *SendMessage) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendMessage {
	message.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return message
}
