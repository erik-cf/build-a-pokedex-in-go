package repl

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandExplore(c *state.PokedexConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("At least one zone must be passed")
	}
	area := args[0]
	fmt.Printf("Exploring %s...\n", area)
	locationArea, err := c.Client.GetSingleLocationArea(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}

	return nil
}
