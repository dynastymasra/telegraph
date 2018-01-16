package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSetWebHook_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestSetWebHook_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestSetWebHook_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusBadRequest).XML(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: bad webhook: HTTPS url must be provided for webhook"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.Equal(t, http.StatusBadRequest, status.StatusCode)
	assert.Error(t, err)
}

func TestDeleteWebHook_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was deleted"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteWebHook_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestDeleteWebHook_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusUnauthorized, status.StatusCode)
	assert.Error(t, err)
}

func TestSendChatAction_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SendChatAction("id", "action").Commit()

	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestSendChatAction_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SendChatAction("id", "action").Commit()

	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestSendChatAction_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SendChatAction("id", "action").Commit()

	assert.Equal(t, http.StatusUnauthorized, status.StatusCode)
	assert.Error(t, err)
}
