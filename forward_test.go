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
