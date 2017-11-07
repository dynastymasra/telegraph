package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetUpdateSuccess(t *testing.T) {
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

	model, status, err := client.GetUpdate().Commit().Parse()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetUpdateSetOffset(t *testing.T) {
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

	model, status, err := client.GetUpdate().SetOffset(5).Commit().Parse()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetUpdateSetLimit(t *testing.T) {
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

	model, status, err := client.GetUpdate().SetLimit(5).Commit().Parse()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetUpdateSetTimeout(t *testing.T) {
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

	model, status, err := client.GetUpdate().SetTimeout(5).Commit().Parse()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetUpdateSetAllowUpdate(t *testing.T) {
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

	model, status, err := client.GetUpdate().SetAllowedUpdates("1", "2", "3").Commit().Parse()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetUpdateError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, status, err := client.GetUpdate().Commit().Parse()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Error(t, err)
}

func TestGetUpdateFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusBadGateway).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, status, err := client.GetUpdate().Commit().Parse()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Error(t, err)
}

func TestGetUpdateFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUpdate, "token")).Reply(http.StatusUnauthorized).XML(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, status, err := client.GetUpdate().Commit().Parse()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Error(t, err)
}
