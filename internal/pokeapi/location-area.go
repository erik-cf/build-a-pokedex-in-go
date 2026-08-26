package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

type PaginatedLocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type LocationArea struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
}

func GetNextLocationArea(c *state.PokedexConfig) (PaginatedResult[PaginatedLocationArea], error) {
	url := c.NextLocation
	return getLocationArea(url, c)
}

func GetPreviousLocationArea(c *state.PokedexConfig) (PaginatedResult[PaginatedLocationArea], error) {
	url := c.PreviousLocation
	if url == "" {
		return PaginatedResult[PaginatedLocationArea]{}, fmt.Errorf("You're on the first page")
	}
	return getLocationArea(url, c)
}

func getLocationArea(url string, c *state.PokedexConfig) (PaginatedResult[PaginatedLocationArea], error) {
	if url == "" {
		return PaginatedResult[PaginatedLocationArea]{}, fmt.Errorf("Invalid url to call PokeAPI: %v", url)
	}

	var body []byte
	v, ok := c.Cache.Get(url)
	if ok {
		fmt.Println("Fetching from cache")
		body = v
	} else {
		fmt.Println("Not fetching from cache")
		response, err := http.Get(url)
		if err != nil {
			return PaginatedResult[PaginatedLocationArea]{}, err
		}
		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if response.StatusCode > 299 {
			return PaginatedResult[PaginatedLocationArea]{}, fmt.Errorf("PokeAPI returned an error: StatusCode %d, Body: %s", response.StatusCode, string(body))
		}
		c.Cache.Add(url, body)
	}

	result, err := NewPaginatedResult[PaginatedLocationArea](body)
	if err != nil {
		return PaginatedResult[PaginatedLocationArea]{}, err
	}

	return *result, nil
}
