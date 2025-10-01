package command

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/mike-jacks/pokedexcli/config"
	pokemonType "github.com/mike-jacks/pokedexcli/internal/types"
)

const (
	minBaseExp = 1
	maxBaseExp = 635

	// Probability clamps (floor/ceiling for success chance)
	pMin = 0.05 // never worse than 5%
	pMax = 0.90 // never better than 90%

	// Curve sharpness: bigger k makes high-baseExp fall off faster
	k = 5.0
)

func calculateCatchProbability(baseExp int, modifiers ...float64) float64 {
	if baseExp < minBaseExp {
		baseExp = minBaseExp
	}
	if baseExp > maxBaseExp {
		baseExp = maxBaseExp
	}

	// Normalize baseExp to x in [0, 1]
	x := float64(baseExp-minBaseExp) / float64(maxBaseExp-minBaseExp)

	// Smooth hardness via logistic centered at 0.5
	// low x -> (easy), high x -> (hard)
	hardness := 1.0 / (1.0 + math.Exp(-k*(x-0.5)))

	// Convert hardness to success probability (invert)
	p := pMin + (pMax-pMin)*(1.0-hardness)

	// Apply any multiplicative modifiers
	for _, m := range modifiers {
		p += m
	}
	// Final clamp
	if p < 0.01 {
		p = 0.01
	} else if p > 0.98 {
		p = 0.98
	}
	return p

}

func tryCapture(rng *rand.Rand, baseExp int, modifiers ...float64) (bool, float64) {
	p := calculateCatchProbability(baseExp, modifiers...)
	return rng.Float64() < p, p
}

func Catch(config *config.Config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("USAGE: catch <pokemon_name>")
		return fmt.Errorf("     catch command requires a single pokemon name")
	}
	pokemonName := args[0]

	pokemonResp, err := config.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(time.Second)
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	success, p := tryCapture(rng, pokemonResp.BaseExperience)
	fmt.Printf("Success chance: %.2f\n", p*100)
	if !success {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	types := make([]string, len(pokemonResp.Types))
	for i, typ := range pokemonResp.Types {
		types[i] = typ.Type.Name
	}
	pokedex := config.Pokedex
	pokemon := pokemonType.Pokemon{
		Name:   pokemonResp.Name,
		Height: pokemonResp.Height,
		Weight: pokemonResp.Weight,
		Stats: pokemonType.PokemonStats{
			HP:             pokemonResp.Stats[0].BaseStat,
			Attack:         pokemonResp.Stats[1].BaseStat,
			Defense:        pokemonResp.Stats[2].BaseStat,
			SpecialAttack:  pokemonResp.Stats[3].BaseStat,
			SpecialDefense: pokemonResp.Stats[4].BaseStat,
			Speed:          pokemonResp.Stats[5].BaseStat,
		},
		Types: types,
	}
	pokedex.AddPokemon(pokemon)
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
