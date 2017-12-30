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

func TestSendVideo_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideo, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 323,
			"from": {
				"id": 23423423,
				"is_bot": true,
				"first_name": "cube",
				"username": "soft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "Cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510304972,
			"video": {
				"duration": 56,
				"width": 640,
				"height": 360,
				"mime_type": "video/mp4",
				"thumb": {
					"file_id": "AAQFAJDSFDSFN2drMyAAR1OUFoDZdeqdcuAAIC",
					"file_size": 320,
					"width": 90,
					"height": 51
				},
				"file_id": "BAADBQADbAAJDDKD68EpVM5CtxU_7nK9Ag",
				"file_size": 4214667
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVideo(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).SetWidth(1000).
		SetHeight(1000).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideo_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideo, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVideo(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).SetWidth(1000).
		SetHeight(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendVideo_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideo, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	message, res, err := client.SendVideo(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).SetWidth(1000).
		SetHeight(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendVoice_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVoice, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 323,
			"from": {
				"id": 23423423,
				"is_bot": true,
				"first_name": "cube",
				"username": "soft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "Cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510304972,
			"voice": {
				"duration": 56,
				"mime_type": "voice/ogg",
				"file_id": "BAADBQADbAAJDDKD68EpVM5CtxU_7nK9Ag",
				"file_size": 4214667
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVoice(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVoice_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVoice, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVoice(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendVoice_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVoice, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVoice(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetCaption("caption").SetDuration(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendVideoNote_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideoNote, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 323,
			"from": {
				"id": 23423423,
				"is_bot": true,
				"first_name": "cube",
				"username": "soft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "Cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510304972,
			"video_note": {
				"duration": 56,
				"length": 123213,
				"file_id": "BAADBQADbAAJDDKD68EpVM5CtxU_7nK9Ag",
				"file_size": 4214667,
				"thumb": {
					"file_id": "AAQFAJDSFDSFN2drMyAAR1OUFoDZdeqdcuAAIC",
					"file_size": 320,
					"width": 90,
					"height": 51
				}
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVideoNote(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLength(1000).SetDuration(1000).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideoNote_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideoNote, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVideoNote(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLength(1000).SetDuration(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendVideoNote_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVideoNote, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendVideoNote(2434234, "./LICENSE", true).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLength(1000).SetDuration(1000).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMediaGroup_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMediaGroup, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": [
			{
				"message_id": 266,
				"from": {
					"id": 123213123,
					"is_bot": true,
					"first_name": "Squrecode",
					"username": "SquarecodeBot"
				},
				"chat": {
					"id": 123123123,
					"first_name": "Dimas",
					"last_name": "Ragil T",
					"username": "dynastymasra",
					"type": "private"
				},
				"date": 1514611687,
				"photo": [
					{
						"file_id": "AgADBAADrZM7G8EXZAe74FQYND4BBC5_iRoABLXsrlGezyfrZyUAAgI",
						"file_size": 1448,
						"width": 90,
						"height": 60
					},
					{
						"file_id": "AgADBAADrZM7G8EXZAe74FQYND4BBC5_iRoABKqrvmmBKVrcaCUAAgI",
						"file_size": 34312,
						"width": 320,
						"height": 213
					},
					{
						"file_id": "AgADBAADrZM7G8EXZAe74FQYND4BBC5_iRoABCxvL8Yd5FkiaSUAAgI",
						"file_size": 176613,
						"width": 800,
						"height": 533
					},
					{
						"file_id": "AgADBAADrZM7G8EXZAe74FQYND4BBC5_iRoABGB0oHcZejCPaiUAAgI",
						"file_size": 418037,
						"width": 1280,
						"height": 853
					},
					{
						"file_id": "AgADBAADrZM7G8EXZAe74FQYND4BBC5_iRoABBYyAbzewQOoayUAAgI",
						"file_size": 1507837,
						"width": 2560,
						"height": 1707
					}
				],
				"caption": "test"
			}
		]
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendMediaGroup(2434234, []telegraph.InputMedia{}).SetDisableNotification(false).
		SetReplyToMessageID(234324234).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMediaGroup_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMediaGroup, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendMediaGroup(2434234, []telegraph.InputMedia{}).SetDisableNotification(false).
		SetReplyToMessageID(234324234).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMediaGroup_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMediaGroup, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendMediaGroup(2434234, []telegraph.InputMedia{}).SetDisableNotification(false).
		SetReplyToMessageID(234324234).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendLocation_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 323,
			"from": {
				"id": 23423423,
				"is_bot": true,
				"first_name": "cube",
				"username": "soft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "Cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510304972,
			"location": {
				"longitude": 123123213.99,
				"latitude": 123213.544
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendLocation(2434234, 12312312.98, 324234324.67).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLivePeriod(60).Commit()

	assert.NotNil(t, message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendLocation_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendLocation, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendLocation(2434234, 12312312.98, 324234324.67).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLivePeriod(60).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendLocation_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendLocation, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message, res, err := client.SendLocation(2434234, 12312312.98, 324234324.67).SetDisableNotification(false).
		SetForceReply(telegraph.ForceReply{}).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{}).SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).
		SetReplyToMessageID(234324234).SetLivePeriod(60).Commit()

	assert.Nil(t, message)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
