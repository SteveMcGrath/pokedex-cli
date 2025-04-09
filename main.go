package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/stevemcgrath/pokedexcli/internal/pokecache"
)

type CliCommand interface {
	Callback(cmd []string) error
	Description() string
	Name() string
}

type Configuration struct {
	cache   *pokecache.Cache
	pokemon map[string]Pokemon
}

var Commands = make(map[string]CliCommand)
var Config = Configuration{pokecache.NewCache(), make(map[string]Pokemon)}

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

func init() {

	ExitCommand(Commands, &Config)
	HelpCommand(Commands, &Config)
	MapCommands(Commands, &Config)
	ExploreCommand(Commands, &Config)
	CatchCommand(Commands, &Config)
	InspectCommand(Commands, &Config)
	PokedexCommand(Commands, &Config)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		words := cleanInput(line)
		if _, ok := Commands[words[0]]; ok {
			if err := Commands[words[0]].Callback(words[1:]); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
