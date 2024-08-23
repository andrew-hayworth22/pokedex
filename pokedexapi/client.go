package pokedexapi

import (
	"net/http"
	"time"

	"github.com/andrew-hayworth22/pokedex/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(cacheInterval),
	}
}
