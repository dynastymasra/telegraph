package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	ChatType  string
	ParseMode string
)

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"

	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML               = "HTML"
)

type (
	Update struct {
		UpdateID          int64    `json:"update_id"`
		Message           *Message `json:"message,omitempty"`
		EditedMessage     *Message `json:"edited_message,omitempty"`
		ChannelPost       *Message `json:"channel_post,omitempty"`
		EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	}

	Message struct {
		MessageID int64       `json:"message_id"`
		Date      int64       `json:"date"`
		Chat      Chat        `json:"chat"`
		From      *From       `json:"from,omitempty"`
		Text      string      `json:"text,omitempty"`
		Photos    []PhotoSize `json:"photo,omitempty"`
		Sticker   *Sticker    `json:"sticker,omitempty"`
	}

	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int64  `json:"width"`
		Height   int64  `json:"height"`
		FileSize int64  `json:"file_size"`
	}

	Sticker struct {
		FileID       string        `json:"file_id"`
		Width        int64         `json:"width"`
		Height       int64         `json:"height"`
		Thumb        *PhotoSize    `json:"thumb,omitempty"`
		Emoji        string        `json:"emoji,omitempty"`
		SetName      string        `json:"set_name,omitempty"`
		FileSize     int64         `json:"file_size,omitempty"`
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	MaskPosition struct {
		Point  string  `json:"point"`
		XShift float64 `json:"x_shift"`
		YShift float64 `json:"y_shift"`
		Scale  float64 `json:"scale"`
	}

	Chat struct {
		ID                     int64    `json:"id"`
		Type                   ChatType `json:"type"`
		Title                  string   `json:"title,omitempty"`
		Username               string   `json:"username,omitempty"`
		FirstName              string   `json:"first_name,omitempty"`
		LastName               string   `json:"last_name,omitempty"`
		AllMemberAdministrator bool     `json:"all_members_are_administrators,omitempty"`
		Description            string   `json:"description,omitempty"`
		InviteLink             string   `json:"invite_link,omitempty"`
	}

	From struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name,omitempty"`
		Username     string `json:"username,omitempty"`
		LanguageCode string `json:"language_code,omitempty"`
	}

	SendMessage struct {
		ChatID          string       `json:"chat_id"`
		Text            string       `json:"text,omitempty"`
		Mode            ParseMode    `json:"parse_mode,omitempty"`
		DisablePreview  bool         `json:"disable_web_page_preview,omitempty"`
		DisNotification bool         `json:"disable_notification,omitempty"`
		ReplyMessageID  int64        `json:"reply_to_message_id,omitempty"`
		Photo           string       `json:"photo,omitempty"`
		Caption         string       `json:"caption,omitempty"`
		ReplyMarkup     *interface{} `json:"reply_markup,omitempty"`
		endpoint        string       `json:"-"`
	}

	ForceReply struct {
		ForceReply bool `json:"force_reply"`
		Selective  bool `json:"selective,omitempty"`
	}
)

/*
Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'

Optional. Use this parameter if you want to force reply from specific users only. Targets:
1) users that are @mentioned in the text of the Message object;
2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
*/
func NewForceReply(selective bool) *ForceReply {
	return &ForceReply{
		ForceReply: true,
		Selective:  selective,
	}
}

// ReplyMessageToID If the message is a reply, ID of the original message
func (message *SendMessage) ReplyMessageToID(id int64) *SendMessage {
	message.ReplyMessageID = id
	return message
}

// DisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *SendMessage) DisableNotification(disable bool) *SendMessage {
	message.DisNotification = disable
	return message
}

// DisableWebPreview Disables link previews for links in this message
func (message *SendMessage) DisableWebPreview(disable bool) *SendMessage {
	message.DisablePreview = disable
	return message
}

// ParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic,
// fixed-width text or inline URLs in your bot's message.
func (message *SendMessage) ParseMode(mode ParseMode) *SendMessage {
	message.Mode = mode
	return message
}

// SetCaption image caption
func (message *SendMessage) SetCaption(caption string) *SendMessage {
	message.Caption = caption
	return message
}

// SetReplyMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *SendMessage) SetReplyMarkup(reply *interface{}) *SendMessage {
	message.ReplyMarkup = reply
	return message
}

// NewTextMessage Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// Text of the message to be sent
func NewTextMessage(chatID, text string) *SendMessage {
	return &SendMessage{
		ChatID:   chatID,
		Text:     text,
		endpoint: EndpointSendMessage,
	}
}

// NewPhotoMessage build new photo message
func NewPhotoMessage(chatID, photoURL string) *SendMessage {
	return &SendMessage{
		ChatID:   chatID,
		Photo:    photoURL,
		endpoint: EndpointSendPhoto,
	}
}

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (client *Client) SendMessage(message SendMessage) *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}
