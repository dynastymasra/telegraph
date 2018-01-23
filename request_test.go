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
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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
		SetInlineMessageID("test").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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
		SetInlineMessageID("test").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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
		SetInlineMessageID("test").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

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

func TestSetChatPhoto_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatPhoto(32423423, "./LICENSE").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSetChatPhoto_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetChatPhoto, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatPhoto(32423423, "./LICENSE").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatPhoto_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatPhoto, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatPhoto(32423423, "./LICENSE").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestDeleteChatPhoto_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatPhoto, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": true
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatPhoto(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteChatPhoto_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteChatPhoto, "token")).ParamPresent("chat_id").
		Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatPhoto(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestDeleteChatPhoto_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatPhoto, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatPhoto(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatTitle_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatTitle, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatTitle(32423423, "title").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSetChatTitle_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetChatTitle, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatTitle(32423423, "title").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatTitle_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatTitle, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatTitle(32423423, "title").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatDescription_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatDescription, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatDescription(32423423, "desc").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSetChatDescription_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetChatDescription, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatDescription(32423423, "desc").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatDescription_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatDescription, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatDescription(32423423, "desc").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestPinChatMessage_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPinChatMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PinChatMessage(32423423, 23423423).SetDisableNotification(true).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestPinChatMessage_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointPinChatMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PinChatMessage(32423423, 23423423).SetDisableNotification(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestPinChatMessage_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPinChatMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.PinChatMessage(32423423, 23423423).SetDisableNotification(true).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestUnpinChatMessage_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointUnpinChatMessage, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": true
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnpinChatMessage(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestUnpinChatMessage_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointUnpinChatMessage, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnpinChatMessage(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestUnpinChatMessage_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointUnpinChatMessage, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.UnpinChatMessage(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestLeaveChat_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointLeaveChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": true
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.LeaveChat(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestLeaveChat_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointLeaveChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.LeaveChat(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestLeaveChat_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointLeaveChat, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.LeaveChat(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestGetChatMembersCount_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": 2
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatMembersCount(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetChatMembersCount_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatMembersCount(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestGetChatMembersCount_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetChatMembersCount, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.GetChatMembersCount(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSetChatStickerSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatStickerSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatStickerSet(32423423, "name").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSetChatStickerSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetChatStickerSet, "token")).Reply(http.StatusOK).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatStickerSet(32423423, "name").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
}

func TestSetChatStickerSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatStickerSet, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetChatStickerSet(32423423, "name").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
}

func TestDeleteChatStickerSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatStickerSet, "token")).ParamPresent("chat_id").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": true
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatStickerSet(32423423).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteChatStickerSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteChatStickerSet, "token")).ParamPresent("chat_id").
		Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatStickerSet(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteChatStickerSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatStickerSet, "token")).ParamPresent("chat_id").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteChatStickerSet(32423423).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestAnswerCallbackQuery_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAnswerCallbackQuery, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerCallbackQuery("23434234").SetText("text").SetShowAlert(true).
		SetURL("https://www.cubesoft.co.id").SetCacheTime(123123123).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestAnswerCallbackQuery_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointAnswerCallbackQuery, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerCallbackQuery("23434234").SetText("text").SetShowAlert(true).
		SetURL("https://www.cubesoft.co.id").SetCacheTime(123123123).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestAnswerCallbackQuery_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAnswerCallbackQuery, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerCallbackQuery("23434234").SetText("text").SetShowAlert(true).
		SetURL("https://www.cubesoft.co.id").SetCacheTime(123123123).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageText_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageText, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageText("text").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetParseMode("HTML").SetDisableWebPagePreview(true).
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageText_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageText, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageText("text").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetParseMode("HTML").SetDisableWebPagePreview(true).
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageText_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageText, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageText("text").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetParseMode("HTML").SetDisableWebPagePreview(true).
		SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageCaption_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageCaption, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageCaption("caption").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageCaption_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageCaption, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageCaption("caption").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageCaption_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageCaption, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageCaption("caption").SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageReplyMarkup_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageReplyMarkup, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageReplyMarkup().SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestEditMessageReplyMarkup_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointEditMessageReplyMarkup, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageReplyMarkup().SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestEditMessageReplyMarkup_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointEditMessageReplyMarkup, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.EditMessageReplyMarkup().SetChatID(1312312).SetMessageID(2323423).
		SetInlineMessageID("inline").SetReplyMarkup([][]telegraph.InlineKeyboardButton{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteMessage_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteMessage, "token")).ParamPresent("chat_id").
		ParamPresent("message_id").Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": true
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteMessage(23223, 232344).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteMessage_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteMessage, "token")).ParamPresent("chat_id").
		ParamPresent("message_id").Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteMessage(23223, 232344).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteMessage_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteMessage, "token")).ParamPresent("chat_id").
		ParamPresent("message_id").Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: invalid file id"
		}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteMessage(23223, 232344).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestCreateNewStickerSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointCreateNewStickerSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.CreateNewStickerSet(234234234, "name", "title", "./LICENSE", ":)").
		SetContainsMask(true).SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestCreateNewStickerSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointCreateNewStickerSet, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.CreateNewStickerSet(234234234, "name", "title", "./LICENSE", ":)").
		SetContainsMask(true).SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestCreateNewStickerSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointCreateNewStickerSet, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.CreateNewStickerSet(234234234, "name", "title", "./LICENSE", ":)").
		SetContainsMask(true).SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestAddStickerToSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAddStickerToSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AddStickerToSet(13123123, "name", "./LICENSE", "emojis").
		SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestAddStickerToSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointAddStickerToSet, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AddStickerToSet(13123123, "name", "./LICENSE", "emojis").
		SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestAddStickerToSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAddStickerToSet, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AddStickerToSet(13123123, "name", "./LICENSE", "emojis").
		SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSetStickerPositionInSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetStickerPositionInSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetStickerPositionInSet("sticker", 10).Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSetStickerPositionInSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSetStickerPositionInSet, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetStickerPositionInSet("sticker", 10).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSetStickerPositionInSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetStickerPositionInSet, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.SetStickerPositionInSet("sticker", 10).Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteStickerFromSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteStickerFromSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteStickerFromSet("sticker").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestDeleteStickerFromSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointDeleteStickerFromSet, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteStickerFromSet("sticker").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestDeleteStickerFromSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteStickerFromSet, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.DeleteStickerFromSet("sticker").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestAnswerInlineQuery_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAnswerInlineQuery, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerInlineQuery("123123123", telegraph.JSON{}, telegraph.JSON{}).SetCacheTime(10000).
		SetIsPersonal(true).SetNextOffset("offset").SetSwitchPMText("text").SetSwitchPMParameter("param").Commit()

	assert.NotNil(t, body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestAnswerInlineQuery_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointAnswerInlineQuery, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerInlineQuery("123123123", telegraph.JSON{}, telegraph.JSON{}).SetCacheTime(10000).
		SetIsPersonal(true).SetNextOffset("offset").SetSwitchPMText("text").SetSwitchPMParameter("param").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestAnswerInlineQuery_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAnswerInlineQuery, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: invalid file id"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	body, res, err := client.AnswerInlineQuery("123123123", telegraph.JSON{}, telegraph.JSON{}).SetCacheTime(10000).
		SetIsPersonal(true).SetNextOffset("offset").SetSwitchPMText("text").SetSwitchPMParameter("param").Commit()

	assert.Nil(t, body)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
