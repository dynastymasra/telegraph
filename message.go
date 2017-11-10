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
		MessageID             int64              `json:"message_id"`
		From                  *From              `json:"from,omitempty"`
		Date                  int64              `json:"date"`
		Chat                  Chat               `json:"chat"`
		ForwardFrom           *User              `json:"forward_from,omitempty"`
		ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
		ForwardFromMessageID  int64              `json:"forward_from_message_id,omitempty"`
		ForwardSignature      string             `json:"forward_signature,omitempty"`
		ForwardDate           int64              `json:"forward_date,omitempty"`
		ReplyToMessage        string             `json:"reply_to_message,omitempty"`
		EditDate              int64              `json:"edit_date,omitempty"`
		AuthorSignature       string             `json:"author_signature,omitempty"`
		Text                  string             `json:"text,omitempty"`
		Entities              []MessageEntity    `json:"entities,omitempty"`
		CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
		Audio                 *Audio             `json:"audio,omitempty"`
		Document              *Document          `json:"document,omitempty"`
		Game                  *Game              `json:"game,omitempty"`
		Photos                []PhotoSize        `json:"photo,omitempty"`
		Sticker               *Sticker           `json:"sticker,omitempty"`
		Video                 *Video             `json:"video,omitempty"`
		Voice                 *Voice             `json:"voice,omitempty"`
		VideoNote             *VideoNote         `json:"video_note,omitempty"`
		Caption               string             `json:"caption,omitempty"`
		Contact               *Contact           `json:"contact,omitempty"`
		Location              *Location          `json:"location,omitempty"`
		Venue                 *Venue             `json:"venue,omitempty"`
		NewChatMembers        []User             `json:"new_chat_members,omitempty"`
		LeftChatMember        *User              `json:"left_chat_member,omitempty"`
		NewChatTitle          string             `json:"new_chat_title,omitempty"`
		NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`
		DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`
		GroupChatCreated      bool               `json:"group_chat_created,omitempty"`
		SuperGroupChatCreated bool               `json:"supergroup_chat_created,omitempty"`
		ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`
		MigrateToChatID       int64              `json:"migrate_to_chat_id,omitempty"`
		MigrateFromChatID     int64              `json:"migrate_from_chat_id,omitempty"`
		PinnedMessage         *Message           `json:"pinned_message,omitempty"`
		Invoice               *Invoice           `json:"invoice,omitempty"`
		SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
	}

	MessageEntity struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		URL    string `json:"url,omitempty"`
		User   *User  `json:"user,omitempty"`
	}

	Audio struct {
		FileID    string `json:"file_id"`
		Duration  int    `json:"duration"`
		Performer string `json:"performer,omitempty"`
		Title     string `json:"title,omitempty"`
		MimeType  string `json:"mime_type,omitempty"`
		FileSize  int    `json:"file_size,omitempty"`
	}

	Document struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileName string     `json:"file_name,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	Game struct {
		Title        string          `json:"title"`
		Description  string          `json:"description"`
		Photo        []PhotoSize     `json:"photo"`
		Text         string          `json:"text,omitempty"`
		TextEntities []MessageEntity `json:"text_entities,animation"`
		Animation    *Animation      `json:"animation,animation"`
	}

	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		FileSize int    `json:"file_size"`
	}

	Sticker struct {
		FileID       string        `json:"file_id"`
		Width        int           `json:"width"`
		Height       int           `json:"height"`
		Thumb        *PhotoSize    `json:"thumb,omitempty"`
		Emoji        string        `json:"emoji,omitempty"`
		SetName      string        `json:"set_name,omitempty"`
		FileSize     int           `json:"file_size,omitempty"`
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	Video struct {
		FileID   string     `json:"file_id"`
		Width    int64      `json:"width"`
		Height   int64      `json:"height"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	Voice struct {
		FileID   string `json:"file_id"`
		Duration int    `json:"duration"`
		MimeType string `json:"mime_type,omitempty"`
		FileSize int    `json:"file_size,omitempty"`
	}

	VideoNote struct {
		FileID   string     `json:"file_id"`
		Length   int        `json:"length"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	Contact struct {
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name,omitempty"`
		UserID      int64  `json:"user_id,omitempty"`
	}

	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	Venue struct {
		Location     Location `json:"location"`
		Title        string   `json:"title"`
		Address      string   `json:"address"`
		FoursquareID string   `json:"foursquare_id,omitempty"`
	}

	Invoice struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		StartParameter string `json:"start_parameter"`
		Currency       string `json:"currency"`
		TotalAmount    int64  `json:"total_amount"`
	}

	SuccessfulPayment struct {
		Currency                string     `json:"currency"`
		TotalAmount             int64      `json:"total_amount"`
		InvoicePayload          string     `json:"invoice_payload"`
		ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
		OrderInfo               *OrderInfo `json:"order_info,omitempty"`
		TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
		ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
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

	Animation struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileName string     `json:"file_name,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	MaskPosition struct {
		Point  string  `json:"point"`
		XShift float64 `json:"x_shift"`
		YShift float64 `json:"y_shift"`
		Scale  float64 `json:"scale"`
	}

	OrderInfo struct {
		Name            string           `json:"name,omitempty"`
		PhoneNumber     string           `json:"phone_number,omitempty"`
		Email           string           `json:"email,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	}

	ShippingAddress struct {
		CountryCode string `json:"country_code"`
		State       string `json:"state"`
		City        string `json:"city"`
		StreetLine1 string `json:"street_line1"`
		StreetLine2 string `json:"street_line2"`
		PostCode    string `json:"post_code"`
	}

	MessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
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

	SendAudio struct {
		ChatID              string       `json:"chat_id"`
		Audio               string       `json:"audio"`
		Caption             string       `json:"caption,omitempty"`
		Duration            int          `json:"duration,omitempty"`
		Performer           string       `json:"performer,omitempty"`
		Title               string       `json:"title,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}

	SendDocument struct {
		ChatID              string       `json:"chat_id"`
		Document            string       `json:"document"`
		Caption             string       `json:"caption,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}

	SendVideo struct {
		ChatID              string       `json:"chat_id"`
		Video               string       `json:"video"`
		Duration            int          `json:"duration,omitempty"`
		Width               int          `json:"width,omitempty"`
		Height              int          `json:"height,omitempty"`
		Caption             string       `json:"caption,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}

	SendVoice struct {
		ChatID              string       `json:"chat_id"`
		Voice               string       `json:"voice"`
		Caption             string       `json:"caption,omitempty"`
		Duration            int          `json:"duration,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}

	SendVideoNote struct {
		ChatID              string       `json:"chat_id"`
		VideoNote           string       `json:"video_note"`
		Duration            int          `json:"duration,omitempty"`
		Length              int          `json:"length,omitempty"`
		DisableNotification bool         `json:"disable_notification,omitempty"`
		ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         *ReplyMarkup `json:"reply_markup,omitempty"`
		endpoint            string       `json:"-"`
	}
)

/*
NewTextMessage Unique identifier for the target chat or username of the target channel (in the format @channelusername)
Text of the message to be sent
*/
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

/*
NewForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
*/
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

/*
NewPhotoMessage Use this method to send photos. On success, the sent Message is returned.
Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
pass an HTTP URL as a String for Telegram to get a photo from the Internet,
or upload a new photo using multipart/form-data.
*/
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
	}
}

/*
NewDocumentMessage Use this method to send general files. On success, the sent Message is returned.
Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
*/
func NewDocumentMessage(chatId, document string) *SendDocument {
	return &SendDocument{
		ChatID:   chatId,
		Document: document,
		endpoint: EndpointSendDocument,
	}
}

// SetCaption Audio caption, 0-200 characters
func (document *SendDocument) SetCaption(caption string) *SendDocument {
	document.Caption = caption
	return document
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (document *SendDocument) SetDisableNotification(disable bool) *SendDocument {
	document.DisableNotification = disable
	return document
}

// SetReplyToMessageId If the message is a reply, ID of the original message
func (document *SendDocument) SetReplyToMessageId(messageId int64) *SendDocument {
	document.ReplyToMessageID = messageId
	return document
}

// SetForceReply
func (document *SendDocument) SetForceReply(reply ForceReply) *SendDocument {
	document.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return document
}

// SetInlineKeyboardMarkup
func (document *SendDocument) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendDocument {
	document.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return document
}

// SetReplyKeyboardMarkup
func (document *SendDocument) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendDocument {
	document.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return document
}

// SetReplyKeyboardRemove
func (document *SendDocument) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendDocument {
	document.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return document
}

/*
SendDocument Use this method to send general files. On success, the sent Message is returned.
Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
*/
func (client *Client) SendDocument(message SendDocument, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Document, "", "document")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
NewAudioMessage Use this method to send audio files, if you want Telegram clients to display them in the music player.
Your audio must be in the .mp3 format. On success, the sent Message is returned.
Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.
*/
func NewAudioMessage(chatId, audio string) *SendAudio {
	return &SendAudio{
		ChatID:   chatId,
		Audio:    audio,
		endpoint: EndpointSendAudio,
	}
}

// SetCaption Audio caption, 0-200 characters
func (audio *SendAudio) SetCaption(caption string) *SendAudio {
	audio.Caption = caption
	return audio
}

// SetDuration Duration of the audio in seconds
func (audio *SendAudio) SetDuration(duration int) *SendAudio {
	audio.Duration = duration
	return audio
}

// SetPerformer Performer
func (audio *SendAudio) SetPerformer(performer string) *SendAudio {
	audio.Performer = performer
	return audio
}

// SetTitle Track name
func (audio *SendAudio) SetTitle(title string) *SendAudio {
	audio.Title = title
	return audio
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (audio *SendAudio) SetDisableNotification(disable bool) *SendAudio {
	audio.DisableNotification = disable
	return audio
}

// SetReplyToMessageId If the message is a reply, ID of the original message
func (audio *SendAudio) SetReplyToMessageId(messageId int64) *SendAudio {
	audio.ReplyToMessageID = messageId
	return audio
}

// SetForceReply
func (audio *SendAudio) SetForceReply(reply ForceReply) *SendAudio {
	audio.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return audio
}

// SetInlineKeyboardMarkup
func (audio *SendAudio) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendAudio {
	audio.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return audio
}

// SetReplyKeyboardMarkup
func (audio *SendAudio) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendAudio {
	audio.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return audio
}

// SetReplyKeyboardRemove
func (audio *SendAudio) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendAudio {
	audio.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return audio
}

/*
SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player.
Your audio must be in the .mp3 format. On success, the sent Message is returned.
Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.
*/
func (client *Client) SendAudio(message SendAudio, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Audio, "", "audio")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
NewVideoMessage Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
On success, the sent Message is returned.
Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
*/
func NewVideoMessage(chatId, video string) *SendVideo {
	return &SendVideo{
		ChatID:   chatId,
		Video:    video,
		endpoint: EndpointSendVideo,
	}
}

// SetCaption caption, 0-200 characters
func (video *SendVideo) SetCaption(caption string) *SendVideo {
	video.Caption = caption
	return video
}

// SetDuration Duration of the audio in seconds
func (video *SendVideo) SetDuration(duration int) *SendVideo {
	video.Duration = duration
	return video
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (video *SendVideo) SetDisableNotification(disable bool) *SendVideo {
	video.DisableNotification = disable
	return video
}

// SetReplyToMessageId If the message is a reply, ID of the original message
func (video *SendVideo) SetReplyToMessageId(messageId int64) *SendVideo {
	video.ReplyToMessageID = messageId
	return video
}

// SetWidth Video width
func (video *SendVideo) SetWidth(width int) *SendVideo {
	video.Width = width
	return video
}

// SetHeight Video height
func (video *SendVideo) SetHeight(height int) *SendVideo {
	video.Height = height
	return video
}

// SetForceReply
func (video *SendVideo) SetForceReply(reply ForceReply) *SendVideo {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return video
}

// SetInlineKeyboardMarkup
func (video *SendVideo) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendVideo {
	video.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return video
}

// SetReplyKeyboardMarkup
func (video *SendVideo) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendVideo {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return video
}

// SetReplyKeyboardRemove
func (video *SendVideo) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendVideo {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return video
}

/*
SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player.
Your audio must be in the .mp3 format. On success, the sent Message is returned.
Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.
*/
func (client *Client) SendVideo(message SendVideo, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Video, "", "video")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
NewVoiceMessage Use this method to send audio files,
if you want Telegram clients to display the file as a playable voice message.
For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size,
this limit may be changed in the future.
*/
func NewVoiceMessage(chatId, voice string) *SendVoice {
	return &SendVoice{
		ChatID:   chatId,
		Voice:    voice,
		endpoint: EndpointSendVoice,
	}
}

// SetCaption caption, 0-200 characters
func (voice *SendVoice) SetCaption(caption string) *SendVoice {
	voice.Caption = caption
	return voice
}

// SetDuration Duration of the audio in seconds
func (voice *SendVoice) SetDuration(duration int) *SendVoice {
	voice.Duration = duration
	return voice
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (voice *SendVoice) SetDisableNotification(disable bool) *SendVoice {
	voice.DisableNotification = disable
	return voice
}

// SetReplyToMessageId If the message is a reply, ID of the original message
func (voice *SendVoice) SetReplyToMessageId(messageId int64) *SendVoice {
	voice.ReplyToMessageID = messageId
	return voice
}

// SetForceReply
func (voice *SendVoice) SetForceReply(reply ForceReply) *SendVoice {
	voice.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return voice
}

// SetInlineKeyboardMarkup
func (voice *SendVoice) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendVoice {
	voice.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return voice
}

// SetReplyKeyboardMarkup
func (voice *SendVoice) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendVoice {
	voice.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return voice
}

// SetReplyKeyboardRemove
func (voice *SendVoice) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendVoice {
	voice.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return voice
}

/*
SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
On success, the sent Message is returned.
Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
*/
func (client *Client) SendVoice(message SendVoice, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Voice, "", "voice")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
NewVideoNoteMessage As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
Use this method to send video messages. On success, the sent Message is returned.
*/
func NewVideoNoteMessage(chatId, video string) *SendVideoNote {
	return &SendVideoNote{
		ChatID:    chatId,
		VideoNote: video,
		endpoint:  EndpointSendVideoNote,
	}
}

// SetLength Video width and height
func (video *SendVideoNote) SetLength(length int) *SendVideoNote {
	video.Length = length
	return video
}

// SetDuration Duration of the audio in seconds
func (video *SendVideoNote) SetDuration(duration int) *SendVideoNote {
	video.Duration = duration
	return video
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (video *SendVideoNote) SetDisableNotification(disable bool) *SendVideoNote {
	video.DisableNotification = disable
	return video
}

// SetReplyToMessageId If the message is a reply, ID of the original message
func (video *SendVideoNote) SetReplyToMessageId(messageId int64) *SendVideoNote {
	video.ReplyToMessageID = messageId
	return video
}

// SetForceReply
func (video *SendVideoNote) SetForceReply(reply ForceReply) *SendVideoNote {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		nil,
		&reply,
	}
	return video
}

// SetInlineKeyboardMarkup
func (video *SendVideoNote) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *SendVideoNote {
	video.ReplyMarkup = &ReplyMarkup{
		&InlineKeyboardMarkup{
			InlineKeyboard: inline,
		},
		nil,
		nil,
		nil,
	}
	return video
}

// SetReplyKeyboardMarkup
func (video *SendVideoNote) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *SendVideoNote {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		&reply,
		nil,
		nil,
	}
	return video
}

// SetReplyKeyboardRemove
func (video *SendVideoNote) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *SendVideoNote {
	video.ReplyMarkup = &ReplyMarkup{
		nil,
		nil,
		&remove,
		nil,
	}
	return video
}

/*
SendVideoNote As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
Use this method to send video messages. On success, the sent Message is returned.
*/
func (client *Client) SendVideoNote(message SendVideoNote, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.VideoNote, "", "video_note")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
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
