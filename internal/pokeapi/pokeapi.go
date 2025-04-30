package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/t6kke/pokedexcli/internal/pokecache"
)

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

func GetLocationAreaDataWCache(url string, cache *pokecache.Cache) (locationareadata, string, error) {
	var loc_area_data locationareadata
	if len(cache.Data) != 0 {
		cache_data, ok := cache.Get(url)
		if ok {
			//fmt.Println("GETTING DATA FROM CACHE")
			json.Unmarshal(cache_data, &loc_area_data)
			return loc_area_data, string(cache_data), nil
		}
	}

	//fmt.Println("GETTING DATA FROM WEB")
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return  locationareadata{}, "", err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return locationareadata{}, "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return locationareadata{}, "", err
	}

	cache.Add(url, body)

	json.Unmarshal(body, &loc_area_data)
	return loc_area_data, string(body), nil
}
