package main

import (
	"fmt"
)

func commandInspect(args []string) error {
	if len(args) == 0 {
		fmt.Println("no pokemon provided for inspection")
		return nil
	}

	pokemon_name := args[0]

	cli_command := command_options["inspect"]

	_, ok := cli_command.config.pokemons[pokemon_name]
	if !ok {
		fmt.Printf("%s not in inventory\n", pokemon_name)
		return nil
	}

	pokemon_full_data := cli_command.config.pokemons[pokemon_name]

	//output information for end user
	fmt.Printf("Name: %s\n", pokemon_full_data.Name)
	fmt.Printf("Height: %d\n", pokemon_full_data.Height)
	fmt.Printf("Weight: %d\n", pokemon_full_data.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon_full_data.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon_full_data.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
