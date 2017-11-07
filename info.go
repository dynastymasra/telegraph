package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

/*
Use this method to specify a url and receive incoming updates via an outgoing webhook.
Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
containing a JSON-serialized Update. In case of an unsuccessful request,
we will give up after a reasonable amount of attempts. Returns true.
*/
func (client *Client) SetWebHook(webHook string) *VoidResponse {
	body := JSON{
		"url": webHook,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCertificate Upload your public key certificate so that the root certificate in use can be checked.
// See our self-signed guide for details.
func (call *VoidResponse) SetCertificate(path string) *VoidResponse {
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Type(gorequest.TypeMultipart).SendFile(path, "", "certificate"),
	}
}

// SetMaxConnection Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery,
// 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server,
// and higher values to increase your bot’s throughput.
func (call *VoidResponse) SetMaxConnection(conn int) *VoidResponse {
	body := JSON{
		"max_connections": conn,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
SetAllowedUpdates List the types of updates you want your bot to receive.
For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
See Update for a complete list of available update types.
Specify an empty list to receive all updates regardless of type (default).
If not specified, the previous setting will be used.
*/
func (call *VoidResponse) SetAllowedUpdates(allowed ...string) *VoidResponse {
	body := JSON{
		"allowed_updates": allowed,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}
