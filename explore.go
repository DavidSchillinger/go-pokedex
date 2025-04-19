package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	}

	err := printExplore(args[0])
	if err != nil {
		return err
	}

	return nil
}

func printExplore(area string) error {
	fmt.Println("Exploring " + area + "...")

	url := "https://pokeapi.co/api/v2/location-area/" + area

	if data, ok := GlobalCache.Get(url); ok {
		return processEncounterData(data)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status OK, got %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	GlobalCache.Add(url, data)
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
