package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     string
	Previous string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		val, ok := getCommands()[cleanLine[0]]
		if ok {
			err := val.callback()
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("Unknown command")
		}
	}
}
