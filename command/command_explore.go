package command

import (
	"encoding/json"
	"fmt"

	"github.com/mike-jacks/pokedexcli/config"
	"github.com/mike-jacks/pokedexcli/internal/pokeapi"
)

func Explore(config *config.Config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("USAGE: explore <location_name>")
		return fmt.Errorf("     explore command requires a single location name")
	}
	locationName := args[0]

	var specificLocationAreaResp pokeapi.RespSpecificLocationArea
	var err error
	var cacheExists bool

	specificLocationAreaRespCachedBytes, cacheExists := config.Pokecache.Get(pokeapi.BaseURL + "/location-area/" + locationName)
	if cacheExists {
		err = json.Unmarshal(specificLocationAreaRespCachedBytes, &specificLocationAreaResp)
		if err != nil {
			return err
		}
	} else {
		specificLocationAreaResp, err = config.PokeapiClient.ExploreLocation(locationName)
		if err != nil {
			return err
		}
		specificLocationAreaRespCachedBytes, err = json.Marshal(specificLocationAreaResp)
		if err != nil {
			return err
		}
		config.Pokecache.Add(pokeapi.BaseURL+"/location-area/"+locationName, specificLocationAreaRespCachedBytes)
	}
	fmt.Println()
	fmt.Printf("Exploring %s...\n", locationName)
	if len(specificLocationAreaResp.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this location area")
		fmt.Println()
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range specificLocationAreaResp.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
