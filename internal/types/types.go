package types

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  PokemonStats
	Types  []string
}

type PokemonStats struct {
	HP int
	Attack int
	Defense int
	SpecialAttack int
	SpecialDefense int
	Speed int
}