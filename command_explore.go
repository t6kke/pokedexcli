package main

import (
	"fmt"
	"github.com/t6kke/pokedexcli/internal/pokeapi"
)

func commandExplore(args []string) error {
	if len(args) == 0 {
		fmt.Println("Location area not provided, nothing to explore")
		return nil
	}
	cli_command := command_options["explore"]
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	loc_area_data, _, err := pokeapi.GetLocationAreaDataWCache(url, &cli_command.config.api_cache)
	if err != nil {
		//TODO not part of lesson scope but should handle 404 location area not found error for user
		return err
	}

	//output information for end user
	for _, item := range loc_area_data.PokemonEncounters {
		fmt.Println(item.Pokemon.Name)
	}

	return nil
}
