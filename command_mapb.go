package main

import (
	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandMapB(c *state.PokedexConfig, _ ...string) error {
	maps, err := c.Client.GetPreviousLocationArea()
	if err != nil {
		return err
	}

	printMap(*maps, c)
	return nil
}
