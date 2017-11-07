package telegraph

import (
	"github.com/parnurzeal/gorequest"
)

type (
	JSON map[string]interface{}

	PrepareRequest struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)
