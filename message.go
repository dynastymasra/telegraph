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
set upload true if its upload file to telegram.
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
set upload true if its upload file to telegram.
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

set upload true if its upload file to telegram.
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

set upload true if its upload file to telegram.
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

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *MessageResponse) SetDisableNotification(disable bool) *MessageResponse {
	body := JSON{
		"disable_notification": disable,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetReplyToMessageID If the message is a reply, ID of the original message
func (message *MessageResponse) SetReplyToMessageID(id int64) *MessageResponse {
	body := JSON{
		"reply_to_message_id": id,
	}

	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetForceReply Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetForceReply(reply ForceReply) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetInlineKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *MessageResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetReplyKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
}

// SetReplyKeyboardRemove Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *MessageResponse {
	body := JSON{
		"reply_markup": remove,
	}
	return &MessageResponse{
		Client:  message.Client,
		Request: message.Request.Send(body),
	}
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
