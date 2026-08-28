package repl

import (
	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandMap(c *state.PokedexConfig, _ ...string) error {
	maps, err := c.Client.GetNextLocationArea()
	if err != nil {
		return err
	}

	printMap(*maps, c)
	return nil
}
