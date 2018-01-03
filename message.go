package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// MessageResponse struct to handle request and response telegram api
	MessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ArrayMessageResponse struct to handle request and array response telegram api
	ArrayMessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// RawMessageResponse struct to handle request and raw response telegram api
	RawMessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
SendMessage Use this method to send text messages. On success, the sent Message is returned.
*/
func (client *Client) SendMessage(chatId interface{}, text string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"text":    text,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendMessage, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
func (message *MessageResponse) SetParseMode(mode string) *MessageResponse {
	body := JSON{
		"parse_mode": mode,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetDisableWebPagePreview Disables link previews for links in this message
func (message *MessageResponse) SetDisableWebPagePreview(disable bool) *MessageResponse {
	body := JSON{
		"disable_web_page_preview": disable,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
*/
func (client *Client) ForwardMessage(chatId, fromChatId interface{}, messageId int) *MessageResponse {
	body := JSON{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	}

	url := client.baseURL + fmt.Sprintf(EndpointForwardMessage, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendPhoto Use this method to send photos. On success, the sent Message is returned.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendPhoto(chatId interface{}, photo string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"photo":   photo,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendPhoto, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(photo, "", "photo")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player.
Your audio must be in the .mp3 format. On success, the sent Message is returned.
Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendAudio(chatId interface{}, audio string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"audio":   audio,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendAudio, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(audio, "", "audio")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetPerformer Performer
func (message *MessageResponse) SetPerformer(performer string) *MessageResponse {
	body := JSON{
		"performer": performer,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetTitle Track name
func (message *MessageResponse) SetTitle(title string) *MessageResponse {
	body := JSON{
		"title": title,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
SendDocument Use this method to send general files. On success, the sent Message is returned.
Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendDocument(chatId interface{}, document string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id":  chatId,
		"document": document,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendDocument, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(document, "", "document")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendVideo Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
On success, the sent Message is returned.
Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendVideo(chatId interface{}, video string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"video":   video,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendVideo, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(video, "", "video")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetWidth Video width
func (message *MessageResponse) SetWidth(width int) *MessageResponse {
	body := JSON{
		"width": width,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetHeight Video height
func (message *MessageResponse) SetHeight(height int) *MessageResponse {
	body := JSON{
		"height": height,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
On success, the sent Message is returned.
Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendVoice(chatId interface{}, voice string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"voice":   voice,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendVoice, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(voice, "", "voice")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetCaption Photo caption (may also be used when resending photos by file_id), 0-200 characters
func (message *MessageResponse) SetCaption(caption string) *MessageResponse {
	body := JSON{
		"caption": caption,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
SendVideoNote As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
Use this method to send video messages. On success, the sent Message is returned.

Set upload true if its upload file to telegram.
*/
func (client *Client) SendVideoNote(chatId interface{}, videoNote string, upload bool) *MessageResponse {
	body := JSON{
		"chat_id":    chatId,
		"video_note": videoNote,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendVideoNote, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(videoNote, "", "video_note")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetLength Video width and height
func (message *MessageResponse) SetLength(length int) *MessageResponse {
	body := JSON{
		"length": length,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetDuration Duration of the audio in seconds
func (message *MessageResponse) SetDuration(duration int) *MessageResponse {
	body := JSON{
		"duration": duration,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
SendLocation Use this method to send point on the map. On success, the sent Message is returned.
*/
func (client *Client) SendLocation(chatId interface{}, latitude, longitude float64) *MessageResponse {
	body := JSON{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetLivePeriod Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
func (message *MessageResponse) SetLivePeriod(livePeriod int) *MessageResponse {
	body := JSON{
		"livePeriod": livePeriod,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

/*
SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
*/
func (client *Client) SendVenue(chatId interface{}, latitude, longitude float64, title, address string) *MessageResponse {
	body := JSON{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
		"title":     title,
		"address":   address,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendVenue, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetFoursquareID Foursquare identifier of the venue
func (message *MessageResponse) SetFoursquareID(id string) *MessageResponse {
	body := JSON{
		"foursquare_id": id,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *MessageResponse) SetDisableNotification(disable bool) *MessageResponse {
	body := JSON{
		"disable_notification": disable,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyToMessageID If the message is a reply, ID of the original message
func (message *MessageResponse) SetReplyToMessageID(id int64) *MessageResponse {
	body := JSON{
		"reply_to_message_id": id,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetForceReply Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetForceReply(reply ForceReply) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetInlineKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *MessageResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyKeyboardRemove Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *MessageResponse {
	body := JSON{
		"reply_markup": remove,
	}
	message.Request = message.Request.Send(body)

	return message
}

// Commit execute request to telegram
func (message *MessageResponse) Commit() (*Message, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *Message `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = message.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(message.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}

/*
SendMediaGroup Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
Not supported upload file to telegram, use url or file id instead.
*/
func (client *Client) SendMediaGroup(chatId interface{}, media []InputMedia) *ArrayMessageResponse {
	body := JSON{
		"chat_id": chatId,
		"media":   media,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendMediaGroup, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &ArrayMessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *ArrayMessageResponse) SetDisableNotification(disable bool) *ArrayMessageResponse {
	body := JSON{
		"disable_notification": disable,
	}

	return &ArrayMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetReplyToMessageID If the message is a reply, ID of the original message
func (message *ArrayMessageResponse) SetReplyToMessageID(id int64) *ArrayMessageResponse {
	body := JSON{
		"reply_to_message_id": id,
	}

	return &ArrayMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// Commit execute request to telegram
func (message *ArrayMessageResponse) Commit() ([]Message, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result []Message `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = message.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(message.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}

/*
EditMessageLiveLocation Use this method to edit live location messages sent by the bot or via the bot (for inline bots).
A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (client *Client) EditMessageLiveLocation(latitude, longitude float64) *RawMessageResponse {
	body := JSON{
		"latitude":  latitude,
		"longitude": longitude,
	}

	url := client.baseURL + fmt.Sprintf(EndpointEditMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &RawMessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
StopMessageLiveLocation Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires.
On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
*/
func (client *Client) StopMessageLiveLocation() *RawMessageResponse {
	url := client.baseURL + fmt.Sprintf(EndpointStopMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &RawMessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetChatID Required if inline_message_id is not specified.
// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (message *RawMessageResponse) SetChatID(chatId interface{}) *RawMessageResponse {
	body := JSON{
		"chat_id": chatId,
	}

	return &RawMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetMessageID Required if inline_message_id is not specified. Identifier of the sent message
func (message *RawMessageResponse) SetMessageID(messageId int) *RawMessageResponse {
	body := JSON{
		"message_id": messageId,
	}

	return &RawMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetInlineMessageID Required if chat_id and message_id are not specified. Identifier of the inline message
func (message *RawMessageResponse) SetInlineMessageID(inlineMessage string) *RawMessageResponse {
	body := JSON{
		"inline_message_id": inlineMessage,
	}

	return &RawMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetInlineKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *RawMessageResponse) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *RawMessageResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	return &RawMessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// Commit execute request to telegram
func (message *RawMessageResponse) Commit() (*http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
	}{}

	operation := func() error {
		res, body, errs = message.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return MakeHTTPResponse(message.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return res, nil
}
