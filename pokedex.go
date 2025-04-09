package main

import "fmt"

type pokedexCommand struct {
	name        string
	description string
	config      *Configuration
}

func PokedexCommand(cmds map[string]CliCommand, config *Configuration) {
	c := pokedexCommand{
		name:        "pokedex",
		description: "What have we caught?",
		config:      config,
	}
	cmds[c.name] = c
}

func (c pokedexCommand) Name() string        { return c.name }
func (c pokedexCommand) Description() string { return c.description }
func (c pokedexCommand) Callback(cmd []string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range c.config.pokemon {
		fmt.Printf("  - %v\n", name)
	}
	return nil
}
