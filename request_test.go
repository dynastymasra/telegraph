package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSetWebHook_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was set"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestSetWebHook_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestSetWebHook_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusBadRequest).XML(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: bad webhook: HTTPS url must be provided for webhook"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").
		SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, status.StatusCode)
	assert.Error(t, err)
}

func TestDeleteWebHook_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true,
		"description": "Webhook was deleted"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.DeleteWebHook().Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteWebHook_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.DeleteWebHook().Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestDeleteWebHook_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteWebHook, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.DeleteWebHook().Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusUnauthorized, status.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageLiveLocation_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageLiveLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageLiveLocation(12312312.98, 324234324.67).SetChatID(21342321).
		SetMessageID(234234234).SetInlineMessageID("test").
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.NotNil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageLiveLocation_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageLiveLocation, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageLiveLocation(12312312.98, 324234324.67).SetChatID(21342321).
		SetMessageID(234234234).SetInlineMessageID("test").
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageLiveLocation_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageLiveLocation, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageLiveLocation(12312312.98, 324234324.67).SetChatID(21342321).
		SetMessageID(234234234).SetInlineMessageID("test").
		SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendChatAction_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, status, err := client.SendChatAction("id", "action").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestSendChatAction_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, status, err := client.SendChatAction("id", "action").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestSendChatAction_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendChatAction, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")

	body, status, err := client.SendChatAction("id", "action").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusUnauthorized, status.StatusCode)
	assert.Error(t, err)
}

func TestKickChatMember_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointKickChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, status, err := client.KickChatMember("234234234", 123423423).SetUntilDate(2343242342).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, status.StatusCode)
	assert.NoError(t, err)
}

func TestKickChatMember_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointKickChatMember, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, status, err := client.KickChatMember("234234234", 123423423).SetUntilDate(2343242342).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, status.StatusCode)
	assert.Error(t, err)
}

func TestKickChatMember_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointKickChatMember, "token")).Reply(http.StatusUnauthorized).JSON(`{
		"ok": false,
		"error_code": 401,
		"description": "Unauthorized"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, status, err := client.KickChatMember("234234234", 123423423).SetUntilDate(2343242342).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusUnauthorized, status.StatusCode)
	assert.Error(t, err)
}

func TestStopMessageLiveLocation_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointStopMessageLiveLocation, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.StopMessageLiveLocation().SetChatID(21342321).SetMessageID(234234234).
		SetInlineMessageID("test").SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.NotNil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestStopMessageLiveLocation_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointStopMessageLiveLocation, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.StopMessageLiveLocation().SetChatID(21342321).SetMessageID(234234234).
		SetInlineMessageID("test").SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestStopMessageLiveLocation_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointStopMessageLiveLocation, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.StopMessageLiveLocation().SetChatID(21342321).SetMessageID(234234234).
		SetInlineMessageID("test").SetInlineKeyboardMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetContent_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetContent, "token", "path")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetContent("path").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetContent_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetContent, "token", "path")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetContent("path").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetContent_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetContent, "token", "path")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetContent("path").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestUnbanChatMember_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointUnbanChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnbanChatMember("32423423", 23423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}

func TestUnbanChatMember_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointUnbanChatMember, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnbanChatMember("32423423", 23423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestUnbanChatMember_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointUnbanChatMember, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnbanChatMember("32423423", 23423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestRestrictChatMember_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointRestrictChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.RestrictChatMember("32423423", 23423423).SetCanSendMessage(true).
		SetCanSendMediaMessage(true).SetCanSendOtherMessage(true).SetCanAddWebPagePreview(true).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestRestrictChatMember_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointRestrictChatMember, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.RestrictChatMember("32423423", 23423423).SetCanSendMessage(true).
		SetCanSendMediaMessage(true).SetCanSendOtherMessage(true).SetCanAddWebPagePreview(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestRestrictChatMember_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointRestrictChatMember, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.RestrictChatMember("32423423", 23423423).SetCanSendMessage(true).
		SetCanSendMediaMessage(true).SetCanSendOtherMessage(true).SetCanAddWebPagePreview(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestPromoteChatMember_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPromoteChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PromoteChatMember("32423423", 23423423).SetCanChangeInfo(true).
		SetCanPostMessage(true).SetCanEditMessage(true).SetCanDeleteMessage(true).SetCanInviteUser(true).
		SetCanRestrictMember(true).SetCanPinMessage(true).SetCanPromoteMember(true).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestPromoteChatMember_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointPromoteChatMember, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PromoteChatMember("32423423", 23423423).SetCanChangeInfo(true).
		SetCanPostMessage(true).SetCanEditMessage(true).SetCanDeleteMessage(true).SetCanInviteUser(true).
		SetCanRestrictMember(true).SetCanPinMessage(true).SetCanPromoteMember(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestPromoteChatMember_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPromoteChatMember, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PromoteChatMember("32423423", 23423423).SetCanChangeInfo(true).
		SetCanPostMessage(true).SetCanEditMessage(true).SetCanDeleteMessage(true).SetCanInviteUser(true).
		SetCanRestrictMember(true).SetCanPinMessage(true).SetCanPromoteMember(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestExportChatInviteLink_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointExportChatInviteLink, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": "https://invite.link.com"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.ExportChatInviteLink(32423423).Commit()

	assert.NotEmpty(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestExportChatInviteLink_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointExportChatInviteLink, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.ExportChatInviteLink(32423423).Commit()

	assert.Empty(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestExportChatInviteLink_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointExportChatInviteLink, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.ExportChatInviteLink(32423423).Commit()

	assert.Empty(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}
