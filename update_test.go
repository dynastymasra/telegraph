package telegraph_test

import (
	"net/http"
	"testing"

	"telegraph"

	"fmt"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestWebHookParseRequest_Success(t *testing.T) {
	payload := []byte(`{
		"update_id": 651868729,
		"message": {
			"message_id": 19,
			"from": {
				"id": 234234,
				"is_bot": false,
				"first_name": "Dimas",
				"last_name": "Ragil T",
				"username": "dynastymasra",
				"language_code": "en-US"
			},
			"chat": {
				"id": 23423423,
				"first_name": "Dimas",
				"last_name": "Ragil T",
				"username": "dynastymasra",
				"type": "private"
			},
			"date": 1508298329,
			"text": "test text"
		}
	}`)

	message, err := telegraph.WebHookParseRequest(payload)

	assert.NotNil(t, message)
	assert.NoError(t, err)
}

func TestWebHookParseRequest_Failed(t *testing.T) {
	payload := []byte(`<-`)

	message, err := telegraph.WebHookParseRequest(payload)

	assert.Nil(t, message)
	assert.Error(t, err)
}

func TestGetUpdates_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusOK).JSON(`{
	  "ok": true,
	  "result": [
		{
		  "update_id": 1234567890,
		  "message": {
			"message_id": 238,
			"from": {
			  "id": 1234567890,
			  "is_bot": true,
			  "first_name": "Cubesoft",
			  "username": "CubesoftBot"
			},
			"chat": {
			  "id": 1234567890,
			  "first_name": "cube",
			  "last_name": "soft",
			  "username": "cubesoft",
			  "type": "private"
			},
			"date": 1510033702,
			"text": "test"
		  }
		}
	  ]
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUpdates().SetOffset(5).SetLimit(5).SetTimeout(5).SetAllowedUpdates("1", "2").Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetUpdates_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUpdates().SetOffset(5).SetLimit(5).SetTimeout(5).SetAllowedUpdates("1", "2").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetUpdates_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUpdates().Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	assert.Error(t, err)
}
