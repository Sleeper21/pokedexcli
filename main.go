package main

import (
	"time"

	"github.com/Sleeper21/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 1*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
		userPokedex:   map[string]pokeapi.Pokemon{}, // Initialize the pokedex map
	}

	startRepl(cfg)
}
