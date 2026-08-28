package pokeapi

import "fmt"

func (c *PokeApiClient) PrintPokedex() {
	fmt.Println("Your Pokedex:")
	for k, _ := range c.Pokedex {
		fmt.Printf("- %s\n", k)
	}
}
