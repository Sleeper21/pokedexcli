package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		//Read the next entry
		scanner.Scan()
		input := scanner.Text() // Captures the input

		//Clean it
		words, err := cleanInput(input)
		if err != nil {
			fmt.Println(err)
		}
		commandName := words[0]

		// Check if it a supported command
		// loop throught the map and look if the key exists
		// If exists, call the callback
		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		// Executes the corresponding function
		err = cmd.callback()
		if err != nil {
			fmt.Println(err)
		}

		// Check if there was an error getting the read Scan
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading the input:", err)
		}
	}

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func cleanInput(text string) ([]string, error) {
	if strings.TrimSpace(text) == "" {
		return []string{}, errors.New("error. empty input")
	}
	allLower := strings.ToLower(text)
	words := strings.Fields(allLower)
	return words, nil
}
