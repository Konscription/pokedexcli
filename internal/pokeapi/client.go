package pokeapi

import (
	"net/http"
	"time"

	"github.com/Konscription/pokedexcli/internal/pokecache"
)

// Client is a wrapper around the http.Client
type Client struct {
	httpClient http.Client
	pokecache  *pokecache.Cache
}

// NewClient creates a new Client with the given timeout
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: pokecache.NewCache(5 * time.Minute),
	}
}
