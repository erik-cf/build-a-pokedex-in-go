package pokeapi

import "encoding/json"

type PaginatedResult[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func NewPaginatedResult[T any](data []byte) (*PaginatedResult[T], error) {
	result := &PaginatedResult[T]{}
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, err
}
