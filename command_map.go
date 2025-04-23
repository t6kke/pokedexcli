package main

import (
	"fmt"
	"github.com/t6kke/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	fmt.Println("Here I should use pokeapi to call location areas first or next page")
	result, _ := pokeapi.Test("https://pokeapi.co/api/v2/location-area/")
	fmt.Println(result)
	command_options := getCommands()
	cli_command := command_options["map"]
	if cli_command.config == nil {
		fmt.Println("no next or previous page URL set")
		cli_command.config.next_url = "https://pokeapi.co/api/v2/location-area/"      //TODO this assignment is not working
		cli_command.config.previous_url = "https://pokeapi.co/api/v2/location-area/"
	}
	fmt.Println(cli_command.config.next_url)
	fmt.Println(cli_command.config.previous_url)
	return nil
}

func commandMapb() error {
	fmt.Println("Here I should use pokeapi to call location previous page")
	//TODO code here
	return nil
}
