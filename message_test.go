package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSendMessage_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendMessage(2434234, "test").SetDisableNotification(false).
		SetDisableWebPagePreview(false).SetParseMode("HTML").SetForceReply(telegraph.ForceReply{}).
		SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageId(234324234).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessage_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendMessage(2434234, "test").SetDisableNotification(false).
		SetDisableWebPagePreview(false).SetParseMode("HTML").SetForceReply(telegraph.ForceReply{}).
		SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageId(234324234).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessage_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendMessage(2434234, "test").SetDisableNotification(false).
		SetDisableWebPagePreview(false).SetParseMode("HTML").SetForceReply(telegraph.ForceReply{}).
		SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageId(234324234).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
