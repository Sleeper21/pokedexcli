package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous

	// Show all results
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
