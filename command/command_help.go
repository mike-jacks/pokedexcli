package command

import (
	"fmt"

	"github.com/mike-jacks/pokedexcli/config"
)

func Help(config *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := GetCommandsSlice()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	return nil
}
