package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, text []string) error {
	// explore command takes 2 words (explore + location)
	//this: "mt-coronet-1f-route-216-area" for example will count as 1 word
	if len(text) > 2 || len(text) < 2 {
		return errors.New("location not found. Use the map command to find valid location-areas")
	}

	areaName := text[1]
	encounterDetails, err := cfg.pokeapiClient.GetAreaEncounters(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", encounterDetails.Name)
	fmt.Println("Pokemon found:")
	for _, data := range encounterDetails.PokemonEncounters {
		fmt.Println(" -", data.Pokemon.Name)
	}

	return nil
}
