package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMeSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"id": 1234567890,
			"is_bot": true,
			"first_name": "telegraph",
			"username": "telegraph"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetMe().Commit()
	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetMeError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetMe().Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetMeFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusBadGateway).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetMe().Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadGateway, res.StatusCode)
	assert.Error(t, err)
}

func TestGetMeFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetMe().Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	assert.Error(t, err)
}
