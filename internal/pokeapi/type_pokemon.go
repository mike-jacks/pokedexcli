package pokeapi

type RespPokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	IsDefault bool `json:"is_default"`
	Order int `json:"order"`
	Weight int `json:"weight"`
	Abilities []Ability `json:"abilities"`
	Forms []Type `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version Type `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item Type `json:"item"`
		VersionDetails []struct {
			Rarity int `json:"rarity"`
			Version Type `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []struct {
		Move Type `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup Type `json:"version_group"`
			MoveLearnMethod Type `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species Type `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat Type `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type Type `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation Generation `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type Type `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation Generation `json:"generation"`
		Abilities []Ability `json:"abilities"`
	} `json:"past_abilities"`
}

type Ability struct {
	Ability Type `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot int `json:"slot"`
}

type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

