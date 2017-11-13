package telegraph_test

import (
	"telegraph"
	"testing"

	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestMakeHTTPResponseInternalServerError(t *testing.T) {
	agent := gorequest.New()

	res := telegraph.MakeHTTPResponse(agent)

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
