package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	JSON map[string]interface{}
)

/*
SendChatAction Use this method when you need to tell the user that something is happening on the bot's side.
The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
Returns True on success.
*/
func (client *Client) SendChatAction(chatID, action string) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"action":  action,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendChatAction, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
KickChatMember Use this method to kick a user from a group, a supergroup or a channel.
In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc.,
unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) KickChatMember(chatID interface{}, userID int64) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"user_id": userID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointKickChatMember, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetUntilDate Date when the user will be unbanned, unix time.
// If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
func (call *VoidResponse) SetUntilDate(untilDate int64) *VoidResponse {
	body := JSON{
		"until_date": untilDate,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
UnbanChatMember Use this method to unban a previously kicked user in a supergroup or channel.
The user will not return to the group or channel automatically, but will be able to join via link, etc.
The bot must be an administrator for this to work. Returns True on success.
*/
func (client *Client) UnbanChatMember(chatID interface{}, userID int64) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"user_id": userID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointUnbanChatMember, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}
