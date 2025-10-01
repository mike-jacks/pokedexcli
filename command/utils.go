package command

import (
	"strings"

	"github.com/mike-jacks/pokedexcli/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
}

func GetCommandsSlice() []CliCommand {
	commandsSlice := []CliCommand{
		{
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		{
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		{
			Name:        "map",
			Description: "Displays the names of location areas in the Pokemon world",
			Callback:    MapForwards,
		},
		{
			Name:        "mapb",
			Description: "Displays the previous page of location areas in the Pokemon world",
			Callback:    MapBackwards,
		},
		{
			Name:        "explore",
			Description: "Explore the a location area and get a list of Pokemon from that location area",
			Callback:    Explore,
		},
		{
			Name:        "catch",
			Description: "Catch a Pokemon",
			Callback:    Catch,
		},
		{
			Name:        "inspect",
			Description: "Inspect a Pokemon",
			Callback:    Inspect,
		},
		{
			Name:        "pokedex",
			Description: "List all Pokemon in your Pokedex",
			Callback:    Pokedex,
		},
	}
	return commandsSlice
}

func GetCommandsMap() map[string]CliCommand {
	commandsMap := make(map[string]CliCommand)
	for _, command := range GetCommandsSlice() {
		commandsMap[command.Name] = command
	}
	return commandsMap
}

func CleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)
	words := strings.Fields(trimmed)
	return words
}
