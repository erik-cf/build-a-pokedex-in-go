package main

import (
	"github.com/erik-cf/build-a-pokedex-in-go/internal/repl"
	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func main() {
	pokedexConfig := state.NewPokedexConfig()
	repl.StartRepl(&pokedexConfig)
}
