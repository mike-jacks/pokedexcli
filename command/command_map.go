package command

import (
	"encoding/json"
	"fmt"

	"github.com/mike-jacks/pokedexcli/config"
	"github.com/mike-jacks/pokedexcli/internal/pokeapi"
)

func MapForwards(config *config.Config, args ...string) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	if config.NextURL == nil && config.PreviousURL != nil {
		return fmt.Errorf("you're on the last page")
	}

	var cacheExists bool
	var locationsRespBytes []byte
	var locationsResp pokeapi.RespShallowLocations
	var err error

	if config.NextURL != nil {
		locationsRespBytes, cacheExists = config.Pokecache.Get(*config.NextURL)
		if cacheExists {
			err = json.Unmarshal(locationsRespBytes, &locationsResp)
			if err != nil {
				return err
			}
		}
	}
	if !cacheExists {
		locationsResp, err = config.PokeapiClient.ListLocations(config.NextURL)
		if err != nil {
			return err
		}
		locationsRespBytes, err = json.Marshal(locationsResp)
		if err != nil {
			return err
		}
		if config.NextURL != nil {
			config.Pokecache.Add(*config.NextURL, locationsRespBytes)
		}
	}

	config.NextURL = locationsResp.Next
	config.PreviousURL = locationsResp.Previous

	for _, locationArea := range locationsResp.Results {
		fmt.Println(locationArea.Name)
	}

	config.NextURL = locationsResp.Next
	config.PreviousURL = locationsResp.Previous
	return nil
}

func MapBackwards(config *config.Config, args ...string) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	if config.PreviousURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	var cacheExists bool
	var locationRespBytes []byte
	var locationResp pokeapi.RespShallowLocations
	var err error

	if config.PreviousURL != nil {
		locationRespBytes, cacheExists = config.Pokecache.Get(*config.PreviousURL)
		if cacheExists {
			err = json.Unmarshal(locationRespBytes, &locationResp)
			if err != nil {
				return err
			}
		}
	}
	if !cacheExists {
		locationResp, err = config.PokeapiClient.ListLocations(config.PreviousURL)
		if err != nil {
			return err
		}
		locationRespBytes, err = json.Marshal(locationResp)
		if err != nil {
			return err
		}
		if config.PreviousURL != nil {
			config.Pokecache.Add(*config.PreviousURL, locationRespBytes)
		}
	}

	config.NextURL = locationResp.Next
	config.PreviousURL = locationResp.Previous

	for _, locationArea := range locationResp.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
