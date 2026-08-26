package state

import (
	"time"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/pokecache"
)

type PokedexConfig struct {
	NextLocation     string
	PreviousLocation string
	Cache            *pokecache.Cache
}

func NewPokedexConfig() PokedexConfig {
	return PokedexConfig{
		NextLocation:     "https://pokeapi.co/api/v2/location-area/",
		PreviousLocation: "",
		Cache:            pokecache.NewCache(time.Second * 20),
	}
}
