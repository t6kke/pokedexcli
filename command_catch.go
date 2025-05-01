package main

import (
	"fmt"
	"math/rand"
	//"github.com/t6kke/pokedexcli/internal/pokeapi"
)

func commandCatch(args []string) error {
	if len(args) == 0 {
		fmt.Println("no pokemon provided to catch")
		return nil
	}

	fmt.Println("TESTING", args)

	//TODO web queries for catching
	/*cli_command := command_options["explore"]
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	loc_area_data, _, err := pokeapi.GetLocationAreaDataWCache(url, &cli_command.config.api_cache)
	if err != nil {
		//TODO not part of lesson scope but should handle 404 pokemon not found
		return err
	}*/

	//output information for end user
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	random_nbr := rand.Int()
	fmt.Println(random_nbr)
	catch_success := false
	if catch_success {
		fmt.Printf("%s was caught!\n", args[0])
	}
	fmt.Printf("%s escaped!\n", args[0])

	return nil
}
