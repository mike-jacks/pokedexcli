package pokedex

import "github.com/mike-jacks/pokedexcli/internal/types"

type Pokedex struct {
	pokemon map[string]types.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{pokemon: make(map[string]types.Pokemon)}
}

func (p *Pokedex) AddPokemon(pokemon types.Pokemon) {
	p.pokemon[pokemon.Name] = pokemon
}

func (p *Pokedex) GetPokemon(name string) (types.Pokemon, bool) {
	pokemon, exists := p.pokemon[name]
	return pokemon, exists
}

func (p *Pokedex) ListPokemon() []string {
	keys := make([]string, 0, len(p.pokemon))
	for name := range p.pokemon {
		keys = append(keys, name)
	}
	return keys
}
