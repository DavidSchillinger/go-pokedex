package main

import (
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

	data, err := Fetch[locationAreaDetailResponse](
		"https://pokeapi.co/api/v2/location-area/" + area,
	)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}
