package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetChatSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChat, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"id": 23423432,
			"first_name": "Dimas",
			"last_name": "Ragil T",
			"username": "dynastymasra",
			"type": "private",
			"photo": {
				"small_file_id": "AQADBKADBqgxG_jQeQQACC4QvjIABDpyUK1bDUxwZeUAAgI",
				"big_file_id": "AQADDQADBqgxG_jQeQQACC4QvjIABIJ0vp2ffevPZ-UAAgI"
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChat("33242342").Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChat, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChat("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChat, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChat("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChat, "token")).Reply(http.StatusNotFound).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChat("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatAdministratorsSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": [{
			"user": {
				"id": 468813201,
				"is_bot": true,
				"first_name": "Squrecode",
				"username": "SquarecodeBot"
			},
			"status": "member"
		}]
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatAdministrators("33242342").Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatAdministratorsError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatAdministrators("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatAdministratorsFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatAdministrators("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatAdministratorsFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).Reply(http.StatusNotFound).XML(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatAdministrators("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMembersCountSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": 2
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	count, res, err := client.GetChatMembersCount("33242342").Commit()

	assert.NotNil(t, count)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatMembersCountError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	count, res, err := client.GetChatMembersCount("33242342").Commit()

	assert.Nil(t, count)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMembersCountFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	count, res, err := client.GetChatMembersCount("33242342").Commit()

	assert.Nil(t, count)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMembersCountFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).Reply(http.StatusNotFound).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	count, res, err := client.GetChatMembersCount("33242342").Commit()

	assert.Nil(t, count)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMemberSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"user": {
				"id": 123213,
				"is_bot": false,
				"first_name": "Dimas",
				"last_name": "Ragil T",
				"username": "dynastymasra",
				"language_code": "en-ID"
			},
			"status": "member"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatMember("33242342", 23432423).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatMemberError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatMember("33242342", 23432423).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMemberFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatMember("33242342", 23432423).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMemberFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).Reply(http.StatusNotFound).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetChatMember("33242342", 23432423).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Error(t, err)
}
