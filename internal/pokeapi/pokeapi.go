package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type locationareas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (locationareas, string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return  locationareas{}, "", err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return locationareas{}, "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return locationareas{}, "", err
	}

	var loc_areas locationareas
	json.Unmarshal(body, &loc_areas)
	return loc_areas, string(body), nil
}
