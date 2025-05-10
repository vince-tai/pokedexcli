package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid location name")
	}

	locName := args[0]
	locationRes, err := cfg.pokeapiClient.GetLocation(locName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationRes.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokeEncounters := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", pokeEncounters.Pokemon.Name)
	}

	return nil
}
