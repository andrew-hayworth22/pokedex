package pokedexapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var LOCATIONS_PAGE_SIZE int = 20

type LocationPageResponse struct {
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}

func (c *Client) GetLocationsPage(page int) ([]Location, error) {
	offset := LOCATIONS_PAGE_SIZE * (page - 1)

	url := fmt.Sprintf("%s/location-area?limit=%d&offset=%d", baseURL, LOCATIONS_PAGE_SIZE, offset)
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
			return nil, errors.New("failed to read data from the response")
		}

		c.cache.Add(url, data)
	}
	var responseJSON LocationPageResponse
	if err := json.Unmarshal(data, &responseJSON); err != nil {
		return nil, fmt.Errorf("failed to convert response to JSON: %s", err.Error())
	}

	return responseJSON.Results, nil
}
