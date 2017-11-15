package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestEditMessageTextSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageText, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageText("text").SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetParseMode(telegraph.ParseModeMarkdown).SetDisableWebPagePreview(true).
		SetReplyMarkup(inline).Commit()

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageTextError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageText, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageText("text").SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetParseMode(telegraph.ParseModeMarkdown).SetDisableWebPagePreview(true).
		SetReplyMarkup(inline).Commit()

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestEditMessageCaptionSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageCaption, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageCaption().SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetCaption("caption").SetReplyMarkup(inline).Commit()

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageCaptionError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageCaption, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageCaption().SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetCaption("caption").SetReplyMarkup(inline).Commit()

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestEditMessageReplyMarkupSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageReplyMarkup, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageReplyMarkup().SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetReplyMarkup(inline).Commit()

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageReplyMarkupError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageReplyMarkup, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	inline := [][]telegraph.InlineKeyboardButton{}
	res, err := client.EditMessageReplyMarkup().SetChatId(1312312).SetMessageId(2323423).
		SetInlineMessageId("inline").SetReplyMarkup(inline).Commit()

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestDeleteMessageSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	res, err := client.DeleteMessage(23223, 232344).Commit()

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteMessageError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	res, err := client.DeleteMessage(23223, 232344).Commit()

	assert.Nil(t, res)
	assert.Error(t, err)
}
