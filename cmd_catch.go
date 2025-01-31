package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func commandCatch(cfg *config, text []string) error {
	// catch command takes 2 words (catch + pokemon)
	// pokemon names always have 1 word .
	// Eg. mr-mime counts as 1 word
	if len(text) != 2 {
		return errors.New("pokemon not found. Use the explore command to find some valid pokemon\nUse the help command to see how to use the explore command")
	}

	pokemonName := text[1]
	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Try to catch it
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)
	baseExperience := pokemonData.BaseExperience

	wasCaptured := throwPokeball(baseExperience)
	if wasCaptured {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		// save pokemon stats into users pokedex
		cfg.userPokedex[pokemonData.Name] = pokemonData

	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil
}

func throwPokeball(exp int) bool {
	// define a difficulty factor which is a constant that determines how hard it is to catch a pokemon.
	// The higher the factor, the easiest will be to catch a pokemon
	// The difficulty factor can be adjusted here
	difficultyFactor := 100.0

	// Get a random number between 0 and 1
	randomValue := rand.Float64()

	// Calculate the catch rate
	// The catch rate will be between 0-1
	// The higher the catch rate, more chances to catch a pokemon
	catchRate := difficultyFactor / (float64(exp) + difficultyFactor)

	waitForPokeball()

	if randomValue <= catchRate {
		return true
	} else {
		return false
	}
}

func waitForPokeball() {
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(2 * time.Second)
}
