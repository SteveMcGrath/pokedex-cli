package main

import "fmt"

type inspectCommand struct {
	name        string
	description string
	config      *Configuration
}

func InspectCommand(cmds map[string]CliCommand, config *Configuration) {
	c := inspectCommand{
		name:        "inspect",
		description: "Inspect a pokemon",
		config:      config,
	}
	cmds[c.name] = c
}

func (c inspectCommand) Name() string        { return c.name }
func (c inspectCommand) Description() string { return c.description }
func (c inspectCommand) Callback(cmd []string) error {
	if len(cmd) != 1 {
		return fmt.Errorf("not a valid inspect command")
	}
	pokemon, ok := c.config.pokemon[cmd[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStats)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}
	return nil
}
