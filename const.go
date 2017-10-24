package telegraph

const (
	Version         = "1.0.0-Beta"
	UserAgentHeader = "User-Agent"
	BaseURL         = "https://api.telegram.org"

	UserAgent = "Telegraph Go SDK"

	EndpointGetMe          = "/bot%v/getMe"
	EndpointSetWebHook     = "/bot%v/setWebhook"
	EndpointGetUpdate      = "/bot%v/getUpdates"
	EndpointDeleteWebHook  = "/bot%v/deleteWebhook"
	EndpointGetWebHookInfo = "/bot%v/getWebhookInfo"
)
