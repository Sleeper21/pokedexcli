package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, text []string) error {
	if len(cfg.userPokedex) == 0 {
		return errors.New("you don't have any Pokemon. Use the catch command to catch a pokemon")
	}
	for pokemon := range cfg.userPokedex {
		fmt.Println(" -", pokemon)
	}
	return nil
}
