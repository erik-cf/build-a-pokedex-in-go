package pokeapi

import (
	"fmt"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

const locationAreaUrl = "https://pokeapi.co/api/v2/location-area/"

type PaginatedLocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationArea struct {
	Id                int                             `json:"id"`
	Name              string                          `json:"name"`
	GameIndex         int                             `json:"game_index"`
	PokemonEncounters []LocationAreaPokemonEncounters `json:"pokemon_encounters"`
}

type LocationAreaPokemonEncounters struct {
	Pokemon MinimalPokemon `json:"pokemon"`
}

func GetNextLocationArea(c *state.PokedexConfig) (*PaginatedResult[PaginatedLocationArea], error) {
	url := c.NextLocation
	return getPaginatedLocationArea(url, c)
}

func GetPreviousLocationArea(c *state.PokedexConfig) (*PaginatedResult[PaginatedLocationArea], error) {
	url := c.PreviousLocation
	if url == "" {
		return nil, fmt.Errorf("You're on the first page")
	}
	return getPaginatedLocationArea(url, c)
}

func GetSingleLocationArea(c *state.PokedexConfig, name string) (*LocationArea, error) {
	if name == "" {
		return nil, fmt.Errorf("Name is required to explore a location area\n")
	}
	return FetchUnmarshalFromApi[LocationArea](c, locationAreaUrl+name)
}

func getPaginatedLocationArea(url string, c *state.PokedexConfig) (*PaginatedResult[PaginatedLocationArea], error) {
	if url == "" {
		return nil, fmt.Errorf("Invalid url to call PokeAPI: %v\n", url)
	}

	return FetchUnmarshalFromApi[PaginatedResult[PaginatedLocationArea]](c, url)
}
