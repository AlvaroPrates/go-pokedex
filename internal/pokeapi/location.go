package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetLocationAreasPayload struct {
	Count         int             `json:"count"`
	Next          string          `json:"next"`
	Previous      *string         `json:"previous"`
	LocationAreas []LocationAreas `json:"results"`
}

type LocationAreas struct {
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
		Name     *string `json:"name"`
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
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas() (GetLocationAreasPayload, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return GetLocationAreasPayload{}, err
	}

	reqBody, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		return GetLocationAreasPayload{}, fmt.Errorf("response failed with status code: %d \nbody: %s", res.StatusCode, reqBody)
	}

	if err != nil {
		return GetLocationAreasPayload{}, err
	}

	var payload GetLocationAreasPayload

	if err := json.Unmarshal(reqBody, &payload); err != nil {
		return GetLocationAreasPayload{}, err
	}

	fmt.Println(payload)

	return payload, nil
}
