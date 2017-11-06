package telegraph_test

import (
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := telegraph.NewClient("token")

	assert.NotNil(t, client)
}

func TestNewClientWithBackOff(t *testing.T) {
	client := telegraph.NewClientWithBackOff("token", telegraph.NewBackOff(10, -1))

	assert.NotNil(t, client)
}
