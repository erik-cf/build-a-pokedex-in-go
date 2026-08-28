package state

import (
	"time"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/pokeapi"
)

type PokedexConfig struct {
	Client *pokeapi.PokeApiClient
}

func NewPokedexConfig() PokedexConfig {
	return PokedexConfig{
		Client: pokeapi.NewPokeApiClient(time.Second * 5),
	}
}
