package telegraph

const (
	Version         = "1.0.0-Development"
	UserAgentHeader = "User-Agent"
	BaseURL         = "https://api.telegram.org"

	UserAgent = "Telegram Go SDK(Telegraph)"

	EndpointGetMe               = "/bot%v/getMe"
	EndpointSetWebHook          = "/bot%v/setWebhook"
	EndpointGetUpdate           = "/bot%v/getUpdates"
	EndpointDeleteWebHook       = "/bot%v/deleteWebhook"
	EndpointGetWebHookInfo      = "/bot%v/getWebhookInfo"
	EndpointGetFile             = "/bot%v/getFile"
	EndpointGetContent          = "/file/bot%v/%v"
	EndpointGetUserProfilePhoto = "/bot%v/getUserProfilePhotos"
	EndpointSendMessage         = "/bot%v/sendMessage"
	EndpointForwardMessage      = "/bot%v/forwardMessage"
	EndpointSendPhoto           = "/bot%v/sendPhoto"
	EndpointSendAudio           = "/bot%v/sendAudio"
	EndpointSendDocument        = "/bot%v/sendDocument"
)
