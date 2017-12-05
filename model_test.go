package telegraph_test

import (
	"net/http"
	"telegraph"
	"testing"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestMakeHTTPResponse_InternalServerError(t *testing.T) {
	agent := gorequest.New()

	res := telegraph.MakeHTTPResponse(agent)

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
