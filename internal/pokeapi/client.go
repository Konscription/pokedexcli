package pokeapi

import (
	"net/http"
	"time"
)

// Client is a wrapper around the http.Client
type Client struct {
	httpClient http.Client
}

// NewClient creates a new Client with the given timeout
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
