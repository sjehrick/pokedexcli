package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, p string) error {
	if p != "" {
		return errors.New("pokedex does not accept arguments")
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
