package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	GameIndex            int                    `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location             Location               `json:"location"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	MinLevel        int    `json:"min_level"`
	MaxLevel        int    `json:"max_level"`
	ConditionValues []any  `json:"condition_values"`
	Chance          int    `json:"chance"`
	Method          Method `json:"method"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon `json:"pokemon"`
	VersionDetails []struct {
		Version          Version            `json:"version"`
		MaxChance        int                `json:"max_chance"`
		EncounterDetails []EncounterDetails `json:"encounter_details"`
	} `json:"version_details"`
}

func getLocationAreas(currentPage int) []string {
	pageSize := 20
	start := currentPage * pageSize

	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/?offset=" + fmt.Sprint(start) + "&limit=" + fmt.Sprint(pageSize))
	if err != nil {
		fmt.Println("Error fetching location areas: ", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return nil
	}
	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return nil
	}
	results, ok := data["results"].([]any)
	if !ok {
		fmt.Println("Error: unexpected data format for results.")
		return nil
	}
	locationAreas := make([]string, 0)
	for _, result := range results {
		areaMap, ok := result.(map[string]any)
		if !ok {
			fmt.Println("Error: unexpected data format for location area.")
			continue
		}
		if name, ok := areaMap["name"].(string); ok {
			locationAreas = append(locationAreas, name)
		}
	}
	if len(locationAreas) == 0 {
		fmt.Println("No location areas found.")
		return locationAreas
	}
	return locationAreas
}
