package main

import (
	"fmt"
	"os"
)

type exitCommand struct {
	name        string
	description string
}

func ExitCommand(cmds map[string]CliCommand, config *Configuration) {
	c := exitCommand{
		name:        "exit",
		description: "Exit the Pokedex",
	}
	cmds[c.name] = c
}

func (c exitCommand) Name() string        { return c.name }
func (c exitCommand) Description() string { return c.description }

func (c exitCommand) Callback(cmd []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
