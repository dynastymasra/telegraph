package telegraph

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff"
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
	Message struct {
		MessageID int64       `json:"message_id"`
		Date      int64       `json:"date"`
		Chat      Chat        `json:"chat"`
		From      *From       `json:"from,omitempty"`
		Text      string      `json:"text,omitempty"`
		Photos    []PhotoSize `json:"photo,omitempty"`
		Sticker   *Sticker    `json:"sticker,omitempty"`
	}

	MessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
		url     string
	}

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

	ForwardMessage struct {
		ChatID              string `json:"chat_id"`
		FromChatID          string `json:"from_chat_id"`
		DisableNotification bool   `json:"disable_notification,omitempty"`
		MessageID           int64  `json:"message_id"`
		endpoint            string `json:"-"`
	}

	SendPhoto struct {
		ChatID              string       `json:"chat_id"`
		Photo               string       `json:"photo"`
		Caption             string       `json:"caption,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
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

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (client *Client) SendMessage(message SendMessage) *MessageResponse {
	url := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// NewForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
func NewForwardMessage(chatID, fromChatID string, messageID int64) *ForwardMessage {
	return &ForwardMessage{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
		endpoint:   EndpointForwardMessage,
	}
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (forward *ForwardMessage) SetDisableNotification(disable bool) *ForwardMessage {
	forward.DisableNotification = disable
	return forward
}

// ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
func (client *Client) ForwardMessage(message ForwardMessage) *MessageResponse {
	url := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

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

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (client *Client) SendPhoto(message SendPhoto, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Photo, "", "photo")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
		url:     message.Photo,
	}
}

// Commit process request send message to telegram
func (message *MessageResponse) Commit() (*Message, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = message.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, makeHTTPResponse(message.Request), err
	}
	return parseMessage(res, body)
}

func parseMessage(res *http.Response, body []byte) (*Message, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *Message `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}
	return model.Result, res, nil
}
