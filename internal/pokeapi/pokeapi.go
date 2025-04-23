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
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Test(url string) (string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return "", err
	}

	var data locationareas //TODO data should be renamed and returned
	json.Unmarshal(body, &data)
	fmt.Println(data.Count)

	//fmt.Printf("%s\n", body)
	return string(body), nil
}
