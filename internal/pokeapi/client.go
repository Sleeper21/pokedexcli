package pokeapi

import (
	"net/http"
	"time"

	"github.com/Sleeper21/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	httpClient     http.Client
	pokeLocalCache pokecache.Cache
}

// New Client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeLocalCache: pokecache.NewCache(cacheInterval),
	}
}
