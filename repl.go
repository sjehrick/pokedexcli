package main

import (
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays twenty locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous twenty locations",
			callback:    commandMapb,
		},
	}
}
