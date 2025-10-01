package command

import (
	"fmt"

	"github.com/mike-jacks/pokedexcli/config"
)

func Inspect(config *config.Config, args ...string) error {
	pokedex := config.Pokedex
	if len(args) != 1 {
		fmt.Println("USAGE: inspect <pokemon_name>")
		return fmt.Errorf("     inspect command requires a single pokemon name")
	}

	pokemonName := args[0]

	_, err := config.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemon, exists := pokedex.GetPokemon(pokemonName)
	if !exists {
		return fmt.Errorf("You have not caught %s yet. %s is not in your Pokedex.", pokemonName, pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  - HP: %d\n", pokemon.Stats.HP)
	fmt.Printf("  - Attack: %d\n", pokemon.Stats.Attack)
	fmt.Printf("  - Defense: %d\n", pokemon.Stats.Defense)
	fmt.Printf("  - Special Attack: %d\n", pokemon.Stats.SpecialAttack)
	fmt.Printf("  - Special Defense: %d\n", pokemon.Stats.SpecialDefense)
	fmt.Printf("  - Speed: %d\n", pokemon.Stats.Speed)
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ)
	}
	fmt.Println()

	return nil
}
