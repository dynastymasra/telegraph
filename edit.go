package telegraph

import (
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// EditMessageResponse struct to handle request and response telegram api
	EditMessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
EditMessageText Use this method to edit text and game messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (client *Client) EditMessageText(text string) *EditMessageResponse {
	body := JSON{
		"text": text,
	}

	url := client.baseURL + fmt.Sprintf(EndpointEditMessageText, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &EditMessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageCaption Use this method to edit captions of messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (client *Client) EditMessageCaption() *EditMessageResponse {
	url := client.baseURL + fmt.Sprintf(EndpointEditMessageCaption, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &EditMessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageReplyMarkup Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (client *Client) EditMessageReplyMarkup() *EditMessageResponse {
	url := client.baseURL + fmt.Sprintf(EndpointEditMessageReplyMarkup, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &EditMessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteMessage Use this method to delete a message, including service messages, with the following limitations:
- A message can only be deleted if it was sent less than 48 hours ago.
- Bots can delete outgoing messages in groups and supergroups.
- Bots granted can_post_messages permissions can delete outgoing messages in channels.
- If the bot is an administrator of a group, it can delete any message there.
- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
Returns True on success.
*/
func (client *Client) DeleteMessage(chatID interface{}, messageID int64) *EditMessageResponse {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteMessage, client.accessToken)
	request := gorequest.New().Get(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v&message_id=%v", chatID, messageID))

	return &EditMessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetChatId Required if inline_message_id is not specified.
// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (edit *EditMessageResponse) SetChatId(chatID interface{}) *EditMessageResponse {
	body := JSON{
		"chat_id": chatID,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetCaption New caption of the message
func (edit *EditMessageResponse) SetCaption(caption string) *EditMessageResponse {
	body := JSON{
		"caption": caption,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetMessageId Required if inline_message_id is not specified. Identifier of the sent message
func (edit *EditMessageResponse) SetMessageId(messageID int64) *EditMessageResponse {
	body := JSON{
		"message_id": messageID,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetInlineMessageId Required if chat_id and message_id are not specified. Identifier of the inline message
func (edit *EditMessageResponse) SetInlineMessageId(inlineMessageID string) *EditMessageResponse {
	body := JSON{
		"inline_message_id": inlineMessageID,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
func (edit *EditMessageResponse) SetParseMode(mode ParseMode) *EditMessageResponse {
	body := JSON{
		"parse_mode": mode,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetDisableWebPagePreview Disables link previews for links in this message
func (edit *EditMessageResponse) SetDisableWebPagePreview(disable bool) *EditMessageResponse {
	body := JSON{
		"disable_web_page_preview": disable,
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// SetReplyMarkup A JSON-serialized object for a new inline keyboard.
func (edit *EditMessageResponse) SetReplyMarkup(inlineKeyboard [][]InlineKeyboardButton) *EditMessageResponse {
	body := JSON{
		"reply_markup": InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		},
	}
	return &EditMessageResponse{
		Client:  edit.Client,
		Request: edit.Request.Send(body),
	}
}

// Commit request to telegram api
func (edit *EditMessageResponse) Commit() (*http.Response, error) {
	var errs []error
	res := &http.Response{}

	operation := func() error {
		res, _, errs = edit.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, edit.Client.expBackOff); err != nil {
		return nil, err
	}
	return res, nil
}
