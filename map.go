package main

import (
	"fmt"
)

type mapConfig struct {
	next string
	prev string
}

type mapCommand struct {
	name        string
	description string
	state       *mapConfig
	config      *Configuration
	is_next     bool
}

func MapCommands(cmds map[string]CliCommand, config *Configuration) {
	state := mapConfig{}
	next := mapCommand{
		name:        "map",
		description: "Displays map locations",
		is_next:     true,
		state:       &state,
		config:      config,
	}
	prev := mapCommand{
		name:        "mapb",
		description: "Displays the previous map page",
		is_next:     false,
		state:       &state,
		config:      config,
	}

	cmds[next.name] = next
	cmds[prev.name] = prev
}

func (c mapCommand) Name() string        { return c.name }
func (c mapCommand) Description() string { return c.description }

func (c mapCommand) Callback(cmd []string) error {
	url := ""
	if c.is_next && c.state.next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else if c.is_next {
		url = c.state.next
	} else if !c.is_next && c.state.prev == "" {
		return fmt.Errorf("no previous page")
	} else {
		url = c.state.prev
	}

	data := struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}{}

	if err := c.config.cache.GetJSON(url, &data); err != nil {
		return err
	}

	c.state.next = data.Next
	c.state.prev = data.Previous

	for _, loc := range data.Results {
		fmt.Printf("- %v\n", loc.Name)
	}

	return nil
}
