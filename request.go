package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	// JSON struct json type
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

/*
RestrictChatMember Use this method to restrict a user in a supergroup.
The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
*/
func (client *Client) RestrictChatMember(chatID interface{}, userID int64) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"user_id": userID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointRestrictChatMember, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCanSendMessages Pass True, if the user can send text messages, contacts, locations and venues
func (call *VoidResponse) SetCanSendMessages(can bool) *VoidResponse {
	body := JSON{
		"can_send_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanSendMediaMessages Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes,
// implies can_send_messages
func (call *VoidResponse) SetCanSendMediaMessages(can bool) *VoidResponse {
	body := JSON{
		"can_send_media_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanSendOtherMessages Pass True, if the user can send animations, games, stickers and use inline bots,
// implies can_send_media_messages
func (call *VoidResponse) SetCanSendOtherMessages(can bool) *VoidResponse {
	body := JSON{
		"can_send_other_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanAddWebPagePreviews Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
func (call *VoidResponse) SetCanAddWebPagePreviews(can bool) *VoidResponse {
	body := JSON{
		"can_add_web_page_previews": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Pass False for all boolean parameters to demote a user. Returns True on success.
*/
func (client *Client) PromoteChatMember(chatID interface{}, userID int64) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"user_id": userID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointPromoteChatMember, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCanChangeInfo Pass True, if the administrator can change chat title, photo and other settings
func (call *VoidResponse) SetCanChangeInfo(can bool) *VoidResponse {
	body := JSON{
		"can_change_info": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanPostMessages Pass True, if the administrator can create channel posts, channels only
func (call *VoidResponse) SetCanPostMessages(can bool) *VoidResponse {
	body := JSON{
		"can_post_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanEditMessages Pass True, if the administrator can edit messages of other users, channels only
func (call *VoidResponse) SetCanEditMessages(can bool) *VoidResponse {
	body := JSON{
		"can_edit_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanDeleteMessages Pass True, if the administrator can delete messages of other users
func (call *VoidResponse) SetCanDeleteMessages(can bool) *VoidResponse {
	body := JSON{
		"can_delete_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanInviteUsers Pass True, if the administrator can invite new users to the chat
func (call *VoidResponse) SetCanInviteUsers(can bool) *VoidResponse {
	body := JSON{
		"can_invite_users": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanRestrictMembers Pass True, if the administrator can restrict, ban or unban chat members
func (call *VoidResponse) SetCanRestrictMembers(can bool) *VoidResponse {
	body := JSON{
		"can_restrict_members": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanPinMessages Pass True, if the administrator can pin messages, supergroups only
func (call *VoidResponse) SetCanPinMessages(can bool) *VoidResponse {
	body := JSON{
		"can_pin_messages": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCanPromoteMembers Pass True,
// if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted,
// directly or indirectly (promoted by administrators that were appointed by him)
func (call *VoidResponse) SetCanPromoteMembers(can bool) *VoidResponse {
	body := JSON{
		"can_promote_members": can,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
ExportChatInviteLink Use this method to export an invite link to a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns exported invite link as String on success.
*/
func (client *Client) ExportChatInviteLink(chatID interface{}) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointExportChatInviteLink, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatPhoto Use this method to set a new profile photo for the chat. Photos can't be changed for private chats.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) SetChatPhoto(chatID interface{}, path string) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetChatPhoto, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeMultipart).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body).
		SendFile(path, "", "photo")

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteChatPhoto Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must b
*/
func (client *Client) DeleteChatPhoto(chatID interface{}) *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteChatPhoto, client.accessToken)
	request := gorequest.New().Get(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatID))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatTitle Use this method to change the title of a chat. Titles can't be changed for private chats.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) SetChatTitle(chatID interface{}, title string) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"title":   title,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetChatTitle, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatDescription Use this method to change the description of a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) SetChatDescription(chatID interface{}) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetChatDescription, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetDescription New chat description, 0-255 characters
func (call *VoidResponse) SetDescription(description string) *VoidResponse {
	body := JSON{
		"description": description,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
PinChatMessage Use this method to pin a message in a supergroup.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) PinChatMessage(chatID interface{}, messageID int64) *VoidResponse {
	body := JSON{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointPinChatMessage, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetDisableNotification Pass True, if it is not necessary to send a notification to all group members about the new pinned message
func (call *VoidResponse) SetDisableNotification(disable bool) *VoidResponse {
	body := JSON{
		"disable_notification": disable,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
UnpinChatMessage Use this method to unpin a message in a supergroup chat.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.
*/
func (client *Client) UnpinChatMessage(chatID interface{}) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointUnpinChatMessage, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
*/
func (client *Client) LeaveChat(chatID interface{}) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointLeaveChat, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatStickerSet Use this method to set a new group sticker set for a supergroup.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
Returns True on success.
*/
func (client *Client) SetChatStickerSet(chatID interface{}, stickerSetName string) *VoidResponse {
	body := JSON{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetChatStickerSet, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
Returns True on success.
*/
func (client *Client) DeleteChatStickerSet(chatID interface{}) *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteChatStickerSet, client.accessToken)
	request := gorequest.New().Get(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatID))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards.
The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
On success, True is returned.
*/
func (client *Client) AnswerCallbackQuery(callbackQueryID string) *VoidResponse {
	body := JSON{
		"callback_query_id": callbackQueryID,
	}

	url := client.baseURL + fmt.Sprintf(EndpointAnswerCallbackQuery, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetText Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
func (call *VoidResponse) SetText(text string) *VoidResponse {
	body := JSON{
		"text": text,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetShowAlert If true, an alert will be shown by the client instead of a notification at the top of the chat screen.
// Defaults to false.
func (call *VoidResponse) SetShowAlert(alert bool) *VoidResponse {
	body := JSON{
		"show_alert": alert,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetUrl URL that will be opened by the user's client.
// If you have created a Game and accepted the conditions via @Botfather,
// specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.
func (call *VoidResponse) SetUrl(url string) *VoidResponse {
	body := JSON{
		"url": url,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// SetCacheTime The maximum amount of time in seconds that the result of the callback query may be cached client-side.
// Telegram apps will support caching starting in version 3.14. Defaults to 0.
func (call *VoidResponse) SetCacheTime(cache int64) *VoidResponse {
	body := JSON{
		"cache_time": cache,
	}
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}
