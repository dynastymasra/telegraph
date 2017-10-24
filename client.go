package telegraph

import (
	"time"

	"github.com/cenkalti/backoff"
)

type Client struct {
	accessToken string
	baseURL     string
	expBackOff  *backoff.ExponentialBackOff
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		baseURL:     BaseURL,
		expBackOff:  NewBackOff(60, -1),
	}
}

// NewClientWithBackOff constructor for client with retry back off
func NewClientWithBackOff(accessToken string, expBackOff *backoff.ExponentialBackOff) *Client {
	return &Client{
		accessToken: accessToken,
		baseURL:     BaseURL,
		expBackOff:  expBackOff,
	}
}

// NewBackOff declare retry exponential back off with max interval time and max elapsed time in second
func NewBackOff(maxInterval, maxElapsedTime int) *backoff.ExponentialBackOff {
	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.MaxElapsedTime = time.Duration(maxElapsedTime) * time.Second
	expBackOff.MaxInterval = time.Duration(maxInterval) * time.Second
	return expBackOff
}
