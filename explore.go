package main

import "fmt"

type urlSubObj struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationAreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod urlSubObj `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int       `json:"rate"`
			Version urlSubObj `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int       `json:"game_index"`
	Id        int       `json:"id"`
	Location  urlSubObj `json:"location"`
	Name      string    `json:"name"`
	Names     []struct {
		Language urlSubObj `json:"language"`
		Name     string    `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        urlSubObj `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int       `json:"chance"`
				ConditionValues []any     `json:"condition_values"`
				MaxLevel        int       `json:"max_level"`
				Method          urlSubObj `json:"method"`
				MinLevel        int       `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int       `json:"max_chance"`
			Version   urlSubObj `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type exploreCommand struct {
	name        string
	description string
	config      *Configuration
}

func ExploreCommand(cmds map[string]CliCommand, config *Configuration) {
	c := exploreCommand{
		name:        "explore",
		description: "Looks for pokemon in the area",
		config:      config,
	}
	cmds[c.name] = c
}

func (c exploreCommand) Name() string        { return c.name }
func (c exploreCommand) Description() string { return c.description }
func (c exploreCommand) Callback(cmd []string) error {
	if len(cmd) != 1 {
		return fmt.Errorf("invalid explore command")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", cmd[0])

	data := locationAreaDetails{}
	if err := c.config.cache.GetJSON(url, &data); err != nil {
		return err
	}
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
