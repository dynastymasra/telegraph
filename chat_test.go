package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetChat_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
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
	body, res, err := client.GetChat(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChat_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChat(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestGetChat_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChat(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatAdministrator_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
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
	body, res, err := client.GetChatAdministrator(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatAdministrator_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatAdministrator(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestGetChatAdministrator_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatAdministrators, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatAdministrator(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMember_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).ParamPresent("chat_id").
		ParamPresent("user_id").Reply(http.StatusOK).JSON(`{
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
	body, res, err := client.GetChatMember(32423423, 23423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatMember_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).ParamPresent("chat_id").
		ParamPresent("user_id").Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatMember(32423423, 23423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestGetChatMember_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMember, "token")).ParamPresent("chat_id").
		ParamPresent("user_id").Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatMember(32423423, 23423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
