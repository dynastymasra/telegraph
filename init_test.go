package telegraph

import (
	"os"
	"testing"

	"github.com/parnurzeal/gorequest"
)

func setUp() {
	gorequest.DisableTransportSwap = true
}

func tearDown() {}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}
