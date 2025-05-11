package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid pokemon name")
	}
	pokeName := args[0]

	if pokemon, ok := cfg.pokedex[pokeName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, pokeStat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", pokeStat.Stat.Name, pokeStat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokeType := range pokemon.Types {
			fmt.Printf(" - %s\n", pokeType.Type.Name)
		}
		return nil
	}

	fmt.Println("you have not caught this pokemon")
	return nil
}
