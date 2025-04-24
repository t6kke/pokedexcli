package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

var command_options map[string]*cliCommand //this is needed to for working with cliCommand as pointers

func startRepl() {
	command_options = getCommands() //this is needed to for working with cliCommand as pointers
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words_slice := cleanInput(text)

		if len(words_slice) == 1 && words_slice[0] == "" {
			continue
		}

		//fmt.Printf("Your command was: %s\n", words_slice[0]) //this was initial test in early lesson
		command_options := getCommands()
		val, ok := command_options[words_slice[0]]
		if ok {
			val.callback()
		} else {
			fmt.Println("unknown command")
		}
	}
}

func cleanInput(text string) []string {
	result_slice := make([]string, 0)

	if text == "" {
		result_slice = append(result_slice, "")
		return result_slice //TODO this probably needs to be handled with error but not currently part of the lesson, function definition needs to be modified to include erros and also
	}

	temp_slice := strings.Fields(text)
	for _, word := range temp_slice {
		result_slice = append(result_slice, strings.ToLower(word))
	}

	return result_slice
}


type Config struct {
	next_url     string
	previous_url string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *Config
}

func getCommands() map[string]*cliCommand {
	return map[string]*cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Displays first or next 20 location areas",
			callback:    commandMap,
			config:      &Config{next_url: "", previous_url: "",},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
			config:      &Config{next_url: "", previous_url: "",},
		},
	}
}


