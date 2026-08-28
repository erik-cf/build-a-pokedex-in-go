package repl

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandCatch(c *state.PokedexConfig, args ...string) error {
	fmt.Println("Command catch")
	if len(args) == 0 {
		return fmt.Errorf("You must pass a pokemon to be able to catch him\n")
	}
	c.Client.CatchPokemon(args[0])
	return nil
}
