package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetFile_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetFile, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"file_id": "AgADBQALBqgxG_jQeQRAHAUL7cXIIy4QvjIABIJ0vp2ffevPZ-UAAgI",
			"file_size": 39421,
			"file_path": "photos/file_65.jpg"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetFile("33242342").Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetFile_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetFile, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetFile("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetFile_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetFile, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	model, res, err := client.GetFile("33242342").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
