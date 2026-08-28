package repl

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandHelp(config *state.PokedexConfig, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
