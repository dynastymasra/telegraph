package telegraph

type (
	SendPhoto struct {
		ChatID              string       `json:"chat_id"`
		Photo               string       `json:"photo"`
		Caption             string       `json:"caption,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}
)

// NewPhotoMessage Use this method to send photos. On success, the sent Message is returned.
// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
// pass an HTTP URL as a String for Telegram to get a photo from the Internet,
// or upload a new photo using multipart/form-data.
func NewPhotoMessage(chatId, photo string) *SendPhoto {
	return &SendPhoto{
		ChatID:   chatId,
		Photo:    photo,
		endpoint: EndpointSendPhoto,
	}
}

// SetCaption Photo caption (may also be used when resending photos by file_id), 0-200 characters
func (photo *SendPhoto) SetCaption(caption string) *SendPhoto {
	photo.Caption = caption
	return photo
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (photo *SendPhoto) SetDisableNotification(disable bool) *SendPhoto {
	photo.DisableNotification = disable
	return photo
}

// SetReplyToMessageId
func (photo *SendPhoto) SetReplyToMessageId(messageId int64) *SendPhoto {
	photo.ReplyToMessageID = messageId
	return photo
}

// SetForceReply
func (photo *SendPhoto) SetForceReply(reply ForceReply) *SendPhoto {
	photo.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return photo
}

// SetInlineKeyboardMarkup
func (photo *SendPhoto) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendPhoto {
	photo.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return photo
}

// SetReplyKeyboardMarkup
func (photo *SendPhoto) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendPhoto {
	photo.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return photo
}

// SetReplyKeyboardRemove
func (photo *SendPhoto) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendPhoto {
	photo.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return photo
}
