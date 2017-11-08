package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

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
	message := telegraph.NewTextMessage("1233456", "test").SetForceReply(reply)
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
