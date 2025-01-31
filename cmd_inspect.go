package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, text []string) error {
	// The inspect command needs to have 2 words --> inspect <pokemon name>
	// Pokemon names have 1 word: --> mr-mime is 1 word
	// Validate entry
	if len(text) != 2 {
		return errors.New("error, inspect command has to have only 2 words, eg. inspect mr-mime")
	}
	// Check if  the pokemon is in the pokedex
	pokemonName := text[1]
	pokemon, exists := cfg.userPokedex[pokemonName]
	if !exists {
		return errors.New("you have not caught that pokemon yet")
	}

	fmt.Printf("Name: %s\nHeight: %d\nStats:\n", pokemon.Name, pokemon.Height)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
