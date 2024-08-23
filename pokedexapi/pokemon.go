package pokedexapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Name              string
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon
}

type Pokemon struct {
	Name           string
	BaseExperience int `json:"base_experience"`
	Height         int
	Weight         int
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
}

func (c *Client) GetPokemonByName(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)
	data, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, errors.New("failed to connect to the Pokedex")
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with code: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, errors.New("failed to read data from response")
		}

		c.cache.Add(url, data)
	}
	var result Pokemon
	if err := json.Unmarshal(data, &result); err != nil {
		return Pokemon{}, fmt.Errorf("failed to convert response to JSON: %s", err.Error())
	}

	return result, nil
}

func (c *Client) GetPokemonByLocation(location string) ([]Pokemon, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, location)
	data, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, errors.New("failed to connect to the Pokedex")
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with code: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("failed to read data from response")
		}

		c.cache.Add(url, data)
	}
	var responseJSON LocationResponse
	if err := json.Unmarshal(data, &responseJSON); err != nil {
		return nil, fmt.Errorf("failed to convert response to JSON: %s", err.Error())
	}

	pokemonMap := map[string]bool{}
	result := []Pokemon{}
	for _, encounter := range responseJSON.PokemonEncounters {
		_, ok := pokemonMap[encounter.Pokemon.Name]
		if !ok {
			pokemonMap[encounter.Pokemon.Name] = true
			result = append(result, encounter.Pokemon)
		}
	}

	return result, nil
}
