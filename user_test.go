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

func TestGetUserProfilePhotos_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUserProfilePhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"total_count": 3,
			"photos": [
				[
					{
						"file_id": "AgADBQADHqgxG_jQeQRAHAUL7cXIIy4QvjIABDpyUK1bDUxwZeUAAgI",
						"file_size": 11160,
						"width": 160,
						"height": 160
					},
					{
						"file_id": "AgAKBQADBqgxG_jQeQRAHAUL7cXIIy4QvjIABEQQfjmV2aXWZuUAAgI",
						"file_size": 30082,
						"width": 320,
						"height": 320
					},
					{
						"file_id": "AgADSQADBqgxG_jQeQRAHAUL7cXIIy4QvjIABIJ0vp2ffevPZ-UAAgI",
						"file_size": 39421,
						"width": 640,
						"height": 640
					}
				]
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUserProfilePhotos(123123213).SetLimit(10).SetOffset(10).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetUserProfilePhotos_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetUserProfilePhoto, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUserProfilePhotos(2312312).Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetUserProfilePhotos_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUserProfilePhoto, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUserProfilePhotos(234234234).Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	assert.Error(t, err)
}
