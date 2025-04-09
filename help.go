package main

import "fmt"

type helpCommand struct {
	name        string
	description string
}

func HelpCommand(cmds map[string]CliCommand, config *Configuration) {
	c := helpCommand{
		name:        "help",
		description: "Displays a help message",
	}
	cmds[c.name] = c
}

func (c helpCommand) Name() string        { return c.name }
func (c helpCommand) Description() string { return c.description }

func (c helpCommand) Callback(cmd []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for k, v := range Commands {
		fmt.Printf("%v: %v\n", k, v.Description())
	}
	return nil
}
