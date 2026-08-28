package pokeapi

import (
	"fmt"
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

func (c *PokeApiClient) GetNextLocationArea() (*PaginatedResult[PaginatedLocationArea], error) {
	url := c.nextLocation
	return c.getPaginatedLocationArea(url)
}

func (c *PokeApiClient) GetPreviousLocationArea() (*PaginatedResult[PaginatedLocationArea], error) {
	url := c.previousLocation
	if url == "" {
		return nil, fmt.Errorf("You're on the first page")
	}
	return c.getPaginatedLocationArea(url)
}

func (c *PokeApiClient) GetSingleLocationArea(name string) (*LocationArea, error) {
	if name == "" {
		return nil, fmt.Errorf("Name is required to explore a location area\n")
	}
	return FetchUnmarshalFromApi[LocationArea](c, locationAreaUrl+name)
}

func (c *PokeApiClient) getPaginatedLocationArea(url string) (*PaginatedResult[PaginatedLocationArea], error) {
	if url == "" {
		return nil, fmt.Errorf("Invalid url to call PokeAPI: %v\n", url)
	}

	mapResult, err := FetchUnmarshalFromApi[PaginatedResult[PaginatedLocationArea]](c, url)
	if err != nil {
		return nil, err
	}

	c.nextLocation = mapResult.Next
	c.previousLocation = mapResult.Previous
	return mapResult, err
}
