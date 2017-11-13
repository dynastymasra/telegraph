package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSendChatActionSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SendChatAction("id", "action").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSendChatActionError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SendChatAction("id", "action").Commit()

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Error(t, err)
}
