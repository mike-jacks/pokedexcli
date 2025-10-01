package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (*RespPokemon, error) {
	url := BaseURL + "/pokemon/" + pokemonName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pokemon '%s' not found", pokemonName)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pokemonResp := &RespPokemon{}
	err = json.Unmarshal(data, pokemonResp)
	if err != nil {
		return nil, err
	}

	return pokemonResp, nil
}
