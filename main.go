package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	cfg.Next = "https://pokeapi.co/api/v2/location-area"
	cfg.Previous = "null"

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		val, ok := getCommands()[cleanLine[0]]
		if ok {
			err := val.callback(cfg)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("Unknown command")
		}
	}
}
