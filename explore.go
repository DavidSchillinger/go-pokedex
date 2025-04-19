package main

import (
	"encoding/json"
	"fmt"
)

type locationAreaDetailResponse struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
}

func CommandExplore(args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide an area to explore!")
		return nil
	}

	area := args[0]

	fmt.Println("Exploring " + area + "...")

	data, err := CachedFetch("https://pokeapi.co/api/v2/location-area/" + area)
	if err != nil {
		return err
	}

	return processEncounterData(data)
}

func processEncounterData(data []byte) error {
	response := locationAreaDetailResponse{}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}
