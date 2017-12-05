package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetWebHookInfo_Success(t *testing.T) {
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

func TestGetWebHookInfo_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetWebHookInfo, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	info, res, err := client.GetWebHookInfo().Commit()

	assert.Nil(t, info)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetWebHookInfo_Failed(t *testing.T) {
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
