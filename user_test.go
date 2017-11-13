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

func TestGetUserProfilePhotosSuccess(t *testing.T) {
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

func TestGetUserProfilePhotosError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetUserProfilePhoto, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUserProfilePhotos(2312312).Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetUserProfilePhotosFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetUserProfilePhoto, "token")).Reply(http.StatusBadGateway).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetUserProfilePhotos(234234).Commit()
	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadGateway, res.StatusCode)
	assert.Error(t, err)
}

func TestGetUserProfilePhotosFailed(t *testing.T) {
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
