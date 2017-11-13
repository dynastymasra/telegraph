package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSendMessageForwardSuccess(t *testing.T) {
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
	message := telegraph.NewForwardMessage("1233456", "1234567890", 1234567890)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageForwardWithDisable(t *testing.T) {
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
			"text": "*test via server*"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "1232435435", 12323434).SetDisableNotification(true)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageForwardError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 123234234)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageForwardFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 123123213)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageForwardFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 234234234)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextSuccess(t *testing.T) {
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
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextWithMarkdown(t *testing.T) {
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
			"text": "*test via server*"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetParseMode(telegraph.ParseModeMarkdown)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextWithHTML(t *testing.T) {
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
	message := telegraph.NewTextMessage("1233456", "test").SetParseMode(telegraph.ParseModeHTML)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextDisableWebHook(t *testing.T) {
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
	message := telegraph.NewTextMessage("1233456", "test").SetDisableWebPagePreview(true)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextDisableNotification(t *testing.T) {
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
	message := telegraph.NewTextMessage("1233456", "test").SetDisableNotification(true)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextReplyId(t *testing.T) {
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
	message := telegraph.NewTextMessage("1233456", "test").SetReplyMessageToId(1234567890)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewTextMessage("1233456", "test").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendPhotoSuccess(t *testing.T) {
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
	message := telegraph.NewPhotoMessage("1233456", "http://www.cubesoft.com/image/test.jpg").SetCaption("test")
	model, res, err := client.SendPhoto(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoDisableNotification(t *testing.T) {
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
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetDisableNotification(true)
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoSetReplyToMessageId(t *testing.T) {
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
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetReplyToMessageId(342412342)
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendAudioSuccess(t *testing.T) {
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
	message := telegraph.NewAudioMessage("1233456", "http://www.cubesoft.com/audio/test.mp3").SetCaption("ok").
		SetDuration(1000).SetPerformer("Cube").SetTitle("soft").SetDisableNotification(true).
		SetReplyToMessageId(123332)
	model, res, err := client.SendAudio(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendAudioReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewAudioMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendAudio(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendDocumentSuccess(t *testing.T) {
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
	message := telegraph.NewDocumentMessage("1233456", "http://www.cubesoft.com/file/test.pdf").SetCaption("ok").
		SetDisableNotification(true).SetReplyToMessageId(123332)
	model, res, err := client.SendDocument(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendDocumentReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewDocumentMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendDocument(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideoSuccess(t *testing.T) {
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
	message := telegraph.NewVideoMessage("1233456", "http://www.cubesoft.com/file/test.pdf").SetCaption("ok").
		SetDisableNotification(true).SetReplyToMessageId(123332).SetDuration(1000).
		SetWidth(1000).SetHeight(1000)
	model, res, err := client.SendVideo(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideoReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewVideoMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendVideo(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVoiceSuccess(t *testing.T) {
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
	message := telegraph.NewVoiceMessage("1233456", "http://www.cubesoft.com/file/test.pdf").SetCaption("ok").
		SetDisableNotification(true).SetReplyToMessageId(123332).SetDuration(1000)
	model, res, err := client.SendVoice(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVoiceReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewVoiceMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendVoice(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideoNoteSuccess(t *testing.T) {
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
	message := telegraph.NewVideoNoteMessage("1233456", "http://www.cubesoft.com/file/test.pdf").SetLength(123123).
		SetDisableNotification(true).SetReplyToMessageId(123332).SetDuration(1000)
	model, res, err := client.SendVideoNote(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVideoNoteReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewVideoNoteMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendVideoNote(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendLocationSuccess(t *testing.T) {
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
	message := telegraph.NewLocationMessage("1233456", 12312312.98, 324234324.67).SetLivePeriod(123123).
		SetDisableNotification(true).SetReplyToMessageId(123332)
	model, res, err := client.SendLocation(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendLocationReplyMarkup(t *testing.T) {
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
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewLocationMessage("1233456", 324234.98, 23423423.56).SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendLocation(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageLiveLocationSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageLiveLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewEditMessageLiveLocation(1233456.34, 32423423.98).SetChatId("23423432").
		SetInlineMessageId("232423432").SetReplyMarkup(inline).SetMessageId(234324)
	res, err := client.EditMessageLiveLocation(*message)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageLiveLocationError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageLiveLocation, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewEditMessageLiveLocation(1233456.34, 32423423.98).SetChatId("23423432").
		SetInlineMessageId("232423432").SetReplyMarkup(inline).SetMessageId(234324)
	res, err := client.EditMessageLiveLocation(*message)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestStopMessageLiveLocationSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointStopMessageLiveLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewStopMessageLiveLocation().SetChatId("23423432").SetInlineMessageId("232423432").
		SetReplyMarkup(inline).SetMessageId(234324)
	res, err := client.StopMessageLiveLocation(*message)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestStopMessageLiveLocationError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointStopMessageLiveLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewStopMessageLiveLocation().SetChatId("23423432").SetInlineMessageId("232423432").
		SetReplyMarkup(inline).SetMessageId(234324)
	res, err := client.StopMessageLiveLocation(*message)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestSendVenueSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendVenue, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 256,
			"from": {
				"id": 234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "byte"
			},
			"chat": {
				"id": 234234324,
				"first_name": "Dimas",
				"last_name": "Ragil T",
				"username": "dynastymasra",
				"type": "private"
			},
			"date": 1510552858,
			"location": {
				"latitude": 234234.998,
				"longitude": 23423423.09
			},
			"venue": {
				"location": {
					"latitude": 234234.099,
					"longitude": -234234.98
				},
				"title": "title",
				"address": "address"
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewVenueMessage("23423423", "title", "address", 32423423.09, 234234.98).
		SetFoursquareId("23423432").SetDisableNotification(true).SetInlineKeyboardMarkup(inline).
		SetForceReply(telegraph.ForceReply{}).SetReplyToMessageId(234324).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendVenue(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendVenueError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendVenue, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewVenueMessage("23423423", "title", "address", 32423423.09, 234234.98).
		SetFoursquareId("23423432").SetDisableNotification(true).
		SetForceReply(telegraph.ForceReply{}).SetReplyToMessageId(234324)
	model, res, err := client.SendVenue(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendContactSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendContact, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 257,
			"from": {
				"id": 23423432,
				"is_bot": true,
				"first_name": "Squrecode",
				"username": "SquarecodeBot"
			},
			"chat": {
				"id": 234324,
				"first_name": "Dimas",
				"last_name": "Ragil T",
				"username": "dynastymasra",
				"type": "private"
			},
			"date": 1510554426,
			"contact": {
				"phone_number": "324234234234",
				"first_name": "Dimas Ragil",
				"last_name": "Triatmaja",
				"user_id": 23432423
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewContactMessage("23423423", "34234234234234234", "Cube").
		SetDisableNotification(true).SetInlineKeyboardMarkup(inline).SetLastName("byte").
		SetForceReply(telegraph.ForceReply{}).SetReplyToMessageId(234324).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendContact(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendContactError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendContact, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewContactMessage("23423423", "34234234234234234", "Cube").
		SetDisableNotification(true).SetInlineKeyboardMarkup(inline).SetLastName("byte").
		SetForceReply(telegraph.ForceReply{}).SetReplyToMessageId(234324).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendContact(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}
