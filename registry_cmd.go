package main

// a struct type that describes a command
type cliCommand struct {
	name        string
	description string
	callback    func() error
}
