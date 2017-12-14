package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMe_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"id": 2432342,
			"is_bot": true,
			"first_name": "Squrecode",
			"username": "SquarecodeBot"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	user, res, err := client.GetMe().Commit()

	assert.NotNil(t, user)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetMe_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	user, res, err := client.GetMe().Commit()

	assert.Nil(t, user)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetMe_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetMe, "token")).Reply(http.StatusNotFound).JSON(`{
		"ok": false,
		"error_code": 404,
		"description": "Not Found: method not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	user, res, err := client.GetMe().Commit()

	assert.Nil(t, user)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}
