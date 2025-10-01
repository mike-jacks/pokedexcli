package config

import (
	"github.com/mike-jacks/pokedexcli/internal/pokeapi"
	"github.com/mike-jacks/pokedexcli/internal/pokecache"
	"github.com/mike-jacks/pokedexcli/internal/pokedex"
)

type Config struct {
	PokeapiClient pokeapi.Client
	NextURL       *string
	PreviousURL   *string
	Pokecache     *pokecache.Cache
	Pokedex       *pokedex.Pokedex
}
