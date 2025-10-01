package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespSpecificLocationArea, error) {
	url := BaseURL + "/location-area/" + locationName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespSpecificLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespSpecificLocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespSpecificLocationArea{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return RespSpecificLocationArea{}, fmt.Errorf("location '%s' not found", locationName)
	}

	specificLocationAreaResp := RespSpecificLocationArea{}

	err = json.Unmarshal(data, &specificLocationAreaResp)
	if err != nil {
		return RespSpecificLocationArea{}, err
	}

	return specificLocationAreaResp, nil
}
