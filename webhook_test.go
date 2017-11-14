package telegraph_test

import (
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebHookParseRequestSuccess(t *testing.T) {
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

func TestWebHookParseRequestFailed(t *testing.T) {
	payload := []byte(`<-`)

	message, err := telegraph.WebHookParseRequest(payload)

	assert.Nil(t, message)
	assert.Error(t, err)
}
