package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, p string) error {
	if p == "" {
		return errors.New("A valid pokemon name must be provided")
	}

	pokemon, exists := cfg.pokedex[p]
	if exists {
		fmt.Printf("Name: %s\n", p)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, name := range pokemon.Types {
			fmt.Printf("  - %s\n", name.Type.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon")
	}

	return nil
}
