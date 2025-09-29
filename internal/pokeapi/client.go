package pokeapi

import (
	"Pokedex/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient(baseURL string, cacheTTL time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    baseURL,
		cache:      pokecache.NewCache(cacheTTL),
	}
}
