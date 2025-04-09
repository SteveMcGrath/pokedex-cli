package main

import (
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version_group"`
			MoveMethodLearned struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStats int `json:"base_stats"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"past_types"`
}

type catchCommand struct {
	name        string
	description string
	config      *Configuration
}

func CatchCommand(cmds map[string]CliCommand, config *Configuration) {
	c := catchCommand{
		name:        "catch",
		description: "Catches a pokemon",
		config:      config,
	}
	cmds[c.name] = c
}

func (c catchCommand) Name() string        { return c.name }
func (c catchCommand) Description() string { return c.description }
func (c catchCommand) Callback(cmd []string) error {
	if len(cmd) != 1 {
		return fmt.Errorf("invalid catch command")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", cmd[0])

	data := Pokemon{}
	if err := c.config.cache.GetJSON(url, &data); err != nil {
		return err
	}
	chance := rand.Intn(data.BaseExperience + 100)
	fmt.Printf("Throwing a Pokeball at %v...\n", cmd[0])
	if chance > data.BaseExperience {
		fmt.Printf("%v was caught!\n", cmd[0])
		c.config.pokemon[cmd[0]] = data
	} else {
		fmt.Printf("%v escaped!\n", cmd[0])
	}
	return nil
}
