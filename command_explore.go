package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, p string) error {
	if p == "" {
		return errors.New("A valid location area must be provided")
	}

	fmt.Println("Exploring " + p + "...")

	pokemonResp, err := cfg.pokeapiClient.ListPokemonEncounters(p)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range pokemonResp.PokemonEncounters {
		fmt.Println("- " + pokemonEncounter.Pokemon.Name)
	}

	return nil
}
