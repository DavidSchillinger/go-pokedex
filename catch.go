package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var CaughtPokemon = map[string]bool{}

type pokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func CommandCatch(args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name of a Pokemon!")
		return nil
	}

	err := printCatch(args[0])
	if err != nil {
		return err
	}

	return nil
}

func printCatch(pokemon string) error {
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	if data, ok := GlobalCache.Get(url); ok {
		return processPokemonData(data)
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
	return processPokemonData(data)
}

func processPokemonData(data []byte) error {
	response := pokemonResponse{}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	fmt.Println(response.Name, "was caught!")
	CaughtPokemon[response.Name] = true

	return nil
}
