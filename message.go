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

// SetReplyToMessageId If the message is a reply, ID of the original message
func (message *MessageResponse) SetReplyToMessageId(id int64) *MessageResponse {
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
