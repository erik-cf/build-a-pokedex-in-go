package main

import (
	"github.com/erik-cf/build-a-pokedex-in-go/internal/pokeapi"
	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandMap(c *state.PokedexConfig) error {
	maps, err := pokeapi.GetNextLocationArea(c)
	if err != nil {
		return err
	}

	printMap(maps, c)
	return nil
}
