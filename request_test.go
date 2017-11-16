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

func TestKickChatMemberSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointKickChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.KickChatMember("234234234", 123423423).SetUntilDate(2343242342).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestUnbanChatMemberSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointUnbanChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.UnbanChatMember("32423423", 23423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestRestrictChatMemberSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointRestrictChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.RestrictChatMember("32423423", 23423423).SetCanSendMessages(true).
		SetCanSendMediaMessages(true).SetCanSendOtherMessages(true).SetCanAddWebPagePreviews(true).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestPromoteChatMemberSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPromoteChatMember, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.PromoteChatMember("32423423", 23423423).SetCanChangeInfo(true).
		SetCanPostMessages(true).SetCanEditMessages(true).SetCanDeleteMessages(true).SetCanInviteUsers(true).
		SetCanRestrictMembers(true).SetCanPinMessages(true).SetCanPromoteMembers(true).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestExportChatInviteLinkSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointExportChatInviteLink, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.ExportChatInviteLink(32423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetChatPhotoSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SetChatPhoto(32423423, "./LICENSE").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestDeleteChatPhotoSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.DeleteChatPhoto(32423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetChatTitleSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatTitle, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SetChatTitle(32423423, "title").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetDescriptionSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatDescription, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SetChatDescription(32423423).SetDescription("description").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestPinChatMessageSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointPinChatMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.PinChatMessage(32423423, 23423423).SetDisableNotification(true).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestUnpinChatMessageSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointUnpinChatMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.UnpinChatMessage(32423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestLeaveChatSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointLeaveChat, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.LeaveChat(32423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetChatStickerSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetChatStickerSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SetChatStickerSet(32423423, "name").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestDeleteChatStickerSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteChatStickerSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.DeleteChatStickerSet(32423423).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestAnswerCallbackQuerySuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAnswerCallbackQuery, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.AnswerCallbackQuery("23434234").SetText("text").SetShowAlert(true).
		SetUrl("https://www.cubesoft.co.id").SetCacheTime(123123123).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestUploadStickerFileSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointUploadStickerFile, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.UploadStickerFile(234234, "./LICENSE").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestCreateNewStickerSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointCreateNewStickerSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.CreateNewStickerSet(234234, "test", "test", "./LICENSE", ":D", true).
		SetContainsMask(true).SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestAddStickerToSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointAddStickerToSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.AddStickerToSet(234234, "test", "./LICENSE", ":D", true).
		SetContainsMask(true).SetMaskPosition(telegraph.MaskPosition{}).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestSetStickerPositionInSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetStickerPositionInSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.SetStickerPositionInSet("test", 2342342).Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestDeleteStickerFromSetSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointDeleteStickerFromSet, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": true
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	status, err := client.DeleteStickerFromSet("test").Commit()

	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}
