package repl

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandInspect(c *state.PokedexConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("You need to pass a pokemon name to be able to inspect it!\n")
	}

	c.Client.InspectPokemon(args[0])
	return nil
}
