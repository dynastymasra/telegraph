package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSetWebHookSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetWebHookSetCertificate(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetWebHookSetMaxConnection(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetMaxConnection(100).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetWebHookSetAllowedUpdates(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").SetAllowedUpdates("1", "2", "3").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetWebHookError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").Commit()

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Error(t, err)
}

func TestSetWebHookFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusBadGateway).XML(``)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").Commit()

	assert.Equal(t, http.StatusBadGateway, status)
	assert.Error(t, err)
}

func TestSetWebHookFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusBadRequest).XML(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: bad webhook: HTTPS url must be provided for webhook"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.SetWebHook("https://www.cubesoft.co.id").Commit()

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err)
}