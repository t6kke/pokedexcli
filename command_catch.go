package main

import (
	"fmt"
	"math/rand"
	"github.com/t6kke/pokedexcli/internal/pokeapi"
)

func commandCatch(args []string) error {
	if len(args) == 0 {
		fmt.Println("no pokemon provided to catch")
		return nil
	}

	pokemon := args[0]

	cli_command := command_options["catch"]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	pokemon_data, _, err := pokeapi.GetPokemonDataWCache(url, &cli_command.config.api_cache)
	if err != nil {
		//TODO not part of lesson scope but should handle 404 pokemon not found
		return err
	}

	pokemon_base_ex := pokemon_data.BaseExperience
	_, ok := cli_command.config.pokemons[pokemon]
	if ok {
		fmt.Printf("%s already in inventory\n", pokemon)
		return nil
	}

	//output information for end user
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	random_nbr := rand.Intn(350)
	//fmt.Println(random_nbr, " --- " ,pokemon_base_ex)
	if random_nbr < pokemon_base_ex {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon)
	cli_command.config.pokemons[pokemon] = pokemon_data

	return nil
}
