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
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageID(234324234).Commit()

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
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageID(234324234).Commit()

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
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).SetReplyToMessageID(234324234).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendPhoto_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendPhoto(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhoto_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendPhoto(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendPhoto_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendPhoto(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestForwardMessage_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusOK).JSON(`{
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
	message, res, err := client.ForwardMessage(2434234, 2434234, 2434234).
		SetDisableNotification(false).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestForwardMessage_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.ForwardMessage(2434234, 2434234, 2434234).
		SetDisableNotification(false).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestForwardMessage_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.ForwardMessage(2434234, 2434234, 2434234).
		SetDisableNotification(false).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendAudio_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendAudio, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 322343,
			"from": {
				"id": 234234324,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 34234234,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510279759,
			"audio": {
				"duration": 162,
				"mime_type": "audio/mpeg",
				"title": "test",
				"performer": "cube",
				"file_id": "NDNDJF949388JF30",
				"file_size": 2668544
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendAudio(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).
		SetPerformer("performer").SetTitle("title").Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendAudio_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendAudio, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendAudio(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).
		SetPerformer("performer").SetTitle("title").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendAudio_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendAudio, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendAudio(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).
		SetPerformer("performer").SetTitle("title").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendDocument_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendDocument, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 3423,
			"from": {
				"id": 324234234,
				"is_bot": true,
				"first_name": "Cube",
				"username": "Cubesoft"
			},
			"chat": {
				"id": 343534,
				"first_name": "Cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510283812,
			"document": {
				"file_name": "test.pdf",
				"mime_type": "application/pdf",
				"file_id": "3FVD44FRDF44-29lVE_cAg",
				"file_size": 239004
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendDocument(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendDocument_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendDocument, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendDocument(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendDocument_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendDocument, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendDocument(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
