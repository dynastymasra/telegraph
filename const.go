package telegraph

const (
	Version         = "1.0.0"
	UserAgentHeader = "User-Agent"
	BaseURL         = "https://api.telegram.org"

	UserAgent = "Telegram Go SDK(Telegraph)"

	EndpointGetMe                   = "/bot%v/getMe"
	EndpointSetWebHook              = "/bot%v/setWebhook"
	EndpointGetUpdate               = "/bot%v/getUpdates"
	EndpointDeleteWebHook           = "/bot%v/deleteWebhook"
	EndpointGetWebHookInfo          = "/bot%v/getWebhookInfo"
	EndpointSendMessage             = "/bot%v/sendMessage"
	EndpointForwardMessage          = "/bot%v/forwardMessage"
	EndpointSendPhoto               = "/bot%v/sendPhoto"
	EndpointSendAudio               = "/bot%v/sendAudio"
	EndpointSendDocument            = "/bot%v/sendDocument"
	EndpointSendVideo               = "/bot%v/sendVideo"
	EndpointSendVoice               = "/bot%v/sendVoice"
	EndpointSendVideoNote           = "/bot%v/sendVideoNote"
	EndpointSendLocation            = "/bot%v/sendLocation"
	EndpointEditMessageLiveLocation = "/bot%v/editMessageLiveLocation"
	EndpointStopMessageLiveLocation = "/bot%v/stopMessageLiveLocation"
	EndpointSendVenue               = "/bot%v/sendVenue"
	EndpointSendContact             = "/bot%v/sendContact"
	EndpointSendChatAction          = "/bot%v/sendChatAction"
	EndpointGetUserProfilePhoto     = "/bot%v/getUserProfilePhotos"
	EndpointGetFile                 = "/bot%v/getFile"
	EndpointGetContent              = "/file/bot%v/%v"
	EndpointKickChatMember          = "/bot%v/kickChatMember"
	EndpointUnbanChatMember         = "/bot%v/unbanChatMember"
	EndpointRestrictChatMember      = "/bot%v/restrictChatMember"
	EndpointPromoteChatMember       = "/bot%v/promoteChatMember"
	EndpointExportChatInviteLink    = "/bot%v/exportChatInviteLink"
	EndpointSetChatPhoto            = "/bot%v/setChatPhoto"
	EndpointDeleteChatPhoto         = "/bot%v/deleteChatPhoto"
	EndpointSetChatTitle            = "/bot%v/setChatTitle"
	EndpointSetChatDescription      = "/bot%v/setChatDescription"
	EndpointPinChatMessage          = "/bot%v/pinChatMessage"
	EndpointUnpinChatMessage        = "/bot%v/unpinChatMessage"
	EndpointLeaveChat               = "/bot%v/leaveChat"
	EndpointGetChat                 = "/bot%v/getChat"
	EndpointGetChatAdministrators   = "/bot%v/getChatAdministrators"
	EndpointGetChatMembersCount     = "/bot%v/getChatMembersCount"
	EndpointGetChatMember           = "/bot%v/getChatMember"
	EndpointSetChatStickerSet       = "/bot%v/setChatStickerSet"
	EndpointDeleteChatStickerSet    = "/bot%v/deleteChatStickerSet"
	EndpointAnswerCallbackQuery     = "/bot%v/answerCallbackQuery"
)
