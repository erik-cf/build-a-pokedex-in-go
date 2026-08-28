package repl

import "github.com/erik-cf/build-a-pokedex-in-go/internal/state"

func commandPokedex(c *state.PokedexConfig, _ ...string) error {
	c.Client.PrintPokedex()
	return nil
}
