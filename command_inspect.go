package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, type_info := range pokemon.Types {
		fmt.Println("  - ", type_info.Type.Name)
	}

	return nil
}
