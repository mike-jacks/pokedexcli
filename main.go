package main

import (
	"time"

	"github.com/mike-jacks/pokedexcli/config"
	"github.com/mike-jacks/pokedexcli/internal/pokeapi"
	"github.com/mike-jacks/pokedexcli/internal/pokecache"
	"github.com/mike-jacks/pokedexcli/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := &config.Config{
		PokeapiClient: pokeClient,
		Pokecache:     pokecache.NewCache(pokeapi.Interval),
		Pokedex:       pokedex.NewPokedex(),
	}
	startRepl(cfg)
}
