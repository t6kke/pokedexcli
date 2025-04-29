package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/t6kke/pokedexcli/internal/pokecache"
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

func GetLocationAreasWCache(url string, cache *pokecache.Cache) (locationareas, string, error) {
	var loc_areas locationareas
	if len(cache.Data) != 0 {
		cache_data, ok := cache.Get(url)
		if ok {
			//fmt.Println("GETTING DATA FROM CACHE")
			json.Unmarshal(cache_data, &loc_areas)
			return loc_areas, string(cache_data), nil
		}
	}

	//fmt.Println("GETTING DATA FROM WEB")
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

	cache.Add(url, body)

	json.Unmarshal(body, &loc_areas)
	return loc_areas, string(body), nil
}
