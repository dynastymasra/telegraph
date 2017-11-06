package telegraph

import (
	"net/http"
)

type (
	Response struct {
		Response *http.Response
		Body     []byte
		Err      error
	}

	ErrorResponse struct {
		OK          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
