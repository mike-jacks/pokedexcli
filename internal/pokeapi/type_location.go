package pokeapi

type RespSpecificLocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel       int `json:"min_level"`
				MaxLevel       int `json:"max_level"`
				ConditionValue struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"condition_value"`
				Chance int `json:"chance"`
				Method struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Order int    `json:"order"`
					Names []struct {
						Name     string `json:"name"`
						Language struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Official bool   `json:"official"`
							ISO639   string `json:"iso639"`
							ISO3166  string `json:"iso3166"`
						} `json:"language"`
					} `json:"names"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
