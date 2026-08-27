package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/erik-cf/build-a-pokedex-in-go/internal/state"
)

type PaginatedResult[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func FetchUnmarshalFromApi[T any](c *state.PokedexConfig, url string) (*T, error) {
	var body []byte
	v, ok := c.Cache.Get(url)
	if ok {
		fmt.Println("Fetching from cache")
		body = v
	} else {
		fmt.Println("Not fetching from cache")
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if response.StatusCode > 299 {
			return nil, fmt.Errorf("PokeAPI returned an error: StatusCode %d, Body: %s", response.StatusCode, string(body))
		}
		c.Cache.Add(url, body)
	}

	return parseJsonData[T](body)
}

func parseJsonData[T any](data []byte) (*T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return &result, err
}
