package main

import (
	"fmt"
)

func commandPokedex(args []string) error {
	cli_command := command_options["pokedex"]
	if len(cli_command.config.pokemons) < 1 {
		fmt.Println("No pokemon in inventory")
		return nil
	}

	//output information for end user
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range cli_command.config.pokemons {
		fmt.Printf("  -%s\n", pokemon.Name)
	}

	return nil
}
