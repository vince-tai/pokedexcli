package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid pokemon name")
	}
	pokeName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)

	pokeRes, err := cfg.pokeapiClient.GetPokemon(pokeName)
	if err != nil {
		return err
	}

	if rand.Intn(400) <= pokeRes.BaseExperience {
		fmt.Printf("%s escaped!\n", pokeRes.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokeRes.Name)
	cfg.pokedex[pokeRes.Name] = pokeRes

	return nil
}
