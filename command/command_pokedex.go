package command

import (
	"fmt"
	"slices"

	"github.com/mike-jacks/pokedexcli/config"
)

func Pokedex(config *config.Config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("pokedex command does not take any arguments")
	}

	fmt.Println()
	fmt.Println("Pokedex:")
	fmt.Println()
	pokedex := config.Pokedex
	pokemonList := pokedex.ListPokemon()
	if len(pokemonList) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}
	slices.Sort(pokemonList)
	for i, pokemon := range pokemonList {
		fmt.Printf("  %d. %s\n", i+1, pokemon)
	}
	fmt.Println()
	return nil
}
