package main

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/pokeapi"
	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

func printMap(mapResult pokeapi.PaginatedResult[pokeapi.PaginatedLocationArea], c *state.PokedexConfig) {
	for _, v := range mapResult.Results {
		fmt.Printf("%s\n", v.Name)
	}

	c.NextLocation = mapResult.Next
	c.PreviousLocation = mapResult.Previous
}
