package pokeapi

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient creates a new PokeAPI client
func NewClient(timeout time.Duration) Client {
	return Client{httpClient: http.Client{Timeout: timeout}}
}
