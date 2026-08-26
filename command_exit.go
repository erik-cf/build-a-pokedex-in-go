package main

import (
	"fmt"
	"os"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func commandExit(config *state.PokedexConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
