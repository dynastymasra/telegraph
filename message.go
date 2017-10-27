package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"
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
		MessageID int64  `json:"message_id"`
		Date      int64  `json:"date"`
		Chat      Chat   `json:"chat"`
		From      *From  `json:"from,omitempty"`
		Text      string `json:"text,omitempty"`
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
		ChatID          string `json:"chat_id"`
		Text            string `json:"text,omitempty"`
		Mode            string `json:"parse_mode,omitempty"`
		DisablePreview  bool   `json:"disable_web_page_preview,omitempty"`
		DisNotification bool   `json:"disable_notification,omitempty"`
		ReplyMessageID  int64  `json:"reply_to_message_id,omitempty"`
		Endpoint        string `json:"-"`
	}
)

// ReplyMessageToID add id message reply
func (message *SendMessage) ReplyMessageToID(id int64) *SendMessage {
	message.ReplyMessageID = id
	return message
}

// DisableNotification add status disable notification
func (message *SendMessage) DisableNotification(disable bool) *SendMessage {
	message.DisNotification = disable
	return message
}

// DisableWebPreview add status disable web preview
func (message *SendMessage) DisableWebPreview(disable bool) *SendMessage {
	message.DisablePreview = disable
	return message
}

// ParseMode add message with parse mode
func (message *SendMessage) ParseMode(mode string) *SendMessage {
	message.Mode = mode
	return message
}

// NewSendTextMessage build new text message
func NewSendTextMessage(chatID, text string) *SendMessage {
	return &SendMessage{
		ChatID:   chatID,
		Text:     text,
		Endpoint: EndpointSendMessage,
	}
}

// SendMessage request send message to telegram
func (client *Client) SendMessage(message SendMessage) *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(message.Endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}
