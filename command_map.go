package main

import (
	"fmt"
	"github.com/t6kke/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	var url string
	cli_command := command_options["map"]
	if cli_command.config.next_url == "" && cli_command.config.previous_url == "" {
		//fmt.Println("First time running the command no URLs set")
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = cli_command.config.next_url
	}

	location_areas, _, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	next_url := location_areas.Next
	previus_url := location_areas.Previous

	//output information for end user
	if previus_url == "" {
		fmt.Println("You're on the first page\n")
	}
	for _, location := range location_areas.Results {
		fmt.Println(location.Name)
	}
	if next_url == "" {
		fmt.Println("\nYou're on the last page")
	}

	// configuring commands struct conf to know next and previous page
	cli_command.config.next_url = next_url
	cli_command.config.previous_url = previus_url
	return nil
}

func commandMapb() error {
	var url string
	cli_command := command_options["map"]

	if cli_command.config.previous_url == "" {
		fmt.Println("you're on the first page, use 'map' command to list first location areas")
		return nil
	} else {
		url = cli_command.config.previous_url
	}

	location_areas, _, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	next_url := location_areas.Next
	previus_url := location_areas.Previous

	//output information for end user
	if previus_url == "" {
		fmt.Println("You're on the first page\n")
	}
	for _, location := range location_areas.Results {
		fmt.Println(location.Name)
	}
	if next_url == "" {
		fmt.Println("\nYou're on the last page")
	}

	// configuring commands struct conf to know next and previous page
	cli_command.config.next_url = next_url
	cli_command.config.previous_url = previus_url
	return nil
}
