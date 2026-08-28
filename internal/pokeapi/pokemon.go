package pokeapi

import (
	"fmt"
	"math/rand"
	"strings"
)

const pokemonEndpoint = "https://pokeapi.co/api/v2/pokemon/"

type MinimalPokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonTypes `json:"types"`
}

type PokemonStats struct {
	Stat     PokemonStat `json:"stat"`
	BaseStat int         `json:"base_stat"`
}

type PokemonStat struct {
	Name string `json:"name"`
}

type PokemonTypes struct {
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
}

func (c *PokeApiClient) CatchPokemon(name string) (*Pokemon, error) {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := FetchUnmarshalFromApi[Pokemon](c, pokemonEndpoint+name)
	if err != nil {
		return nil, err
	}

	r := rand.Float32()
	catchValue := 1.0 / float32(pokemon.BaseExperience) / r
	fmt.Printf("Random number of %f and pokemon with base experience %d. Catch Value is: 1 / %d / %f = %f\n", r, pokemon.BaseExperience, pokemon.BaseExperience, r, catchValue)
	if catchValue > 0.025 {
		fmt.Printf("Adding %s to Pokedex... Congrats!\n", name)
		c.Pokedex[strings.ToLower(name)] = *pokemon
		return pokemon, nil
	}

	return nil, nil
}

func (c *PokeApiClient) InspectPokemon(name string) {
	pokemon, ok := c.Pokedex[strings.ToLower(name)]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf("\t- %s\n", v.Type.Name)
	}
}
