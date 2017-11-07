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

func TestGetWebHookInfoSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetWebHookInfo, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"url": "https://www.cube.com/webhook",
			"has_custom_certificate": false,
			"pending_update_count": 0,
			"max_connections": 100
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	info, res, err := client.GetWebHookInfo().Commit()

	assert.NotNil(t, info)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetWebHookInfoError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetWebHookInfo, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	info, res, err := client.GetWebHookInfo().Commit()

	assert.Nil(t, info)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetWebHookInfoFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetWebHookInfo, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	info, res, err := client.GetWebHookInfo().Commit()

	assert.Nil(t, info)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetWebHookInfoFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetWebHookInfo, "token")).Reply(http.StatusNotFound).JSON(`{
		"ok": false,
		"error_code": 404,
		"description": "Not Found: method not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	info, res, err := client.GetWebHookInfo().Commit()

	assert.Nil(t, info)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteWebHookSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was deleted"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestDeleteWebHookError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Error(t, err)
}

func TestDeleteWebHookFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err)
}

func TestDeleteWebHookFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	status, err := client.DeleteWebHook().Commit()

	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Error(t, err)
}