package pokeapi

import (
	"time"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/pokecache"
)

type PokeApiClient struct {
	Cache            *pokecache.Cache
	nextLocation     string
	previousLocation string
}

func NewPokeApiClient(interval time.Duration) *PokeApiClient {
	return &PokeApiClient{
		Cache:            pokecache.NewCache(time.Second * 20),
		nextLocation:     "https://pokeapi.co/api/v2/location-area/",
		previousLocation: "",
	}
}
