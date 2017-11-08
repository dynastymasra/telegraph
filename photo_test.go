package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSendPhotoSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetCaption("test")
	model, res, err := client.SendPhoto(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoDisableNotification(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetDisableNotification(true)
	model, res, err := client.SendPhoto(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoSetReplyToMessageId(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetReplyToMessageId(342412342)
	model, res, err := client.SendPhoto(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoReplyMarkup(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendPhoto(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}
