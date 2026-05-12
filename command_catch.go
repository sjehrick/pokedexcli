package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, p string) error {
	if p == "" {
		return errors.New("A valid pokemon name must be provided")
	}

	fmt.Println("Throwing a Pokeball at " + p + "...")

	respPokemon, err := cfg.pokeapiClient.ListPokemonStats(p)
	if err != nil {
		return err
	}

	// 675 is based on the ~10% increase above max base experience of Blissey at 608
	randCatch := rand.Intn(675)

	if respPokemon.BaseExperience <= randCatch {
		fmt.Println(p + " was caught!")
		cfg.pokedex[p] = respPokemon
		return nil
	}

	fmt.Println(p + " escaped!")

	return nil
}
