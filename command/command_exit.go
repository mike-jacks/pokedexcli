package command

import (
	"fmt"
	"os"

	"github.com/mike-jacks/pokedexcli/config"
)

func Exit(config *config.Config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("exit command does not take any arguments")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
