package main

import (
	"fmt"
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

	return printCatch(args[0])
}

func printCatch(pokemon string) error {
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")

	data, err := Fetch[pokemonResponse](
		"https://pokeapi.co/api/v2/pokemon/" + pokemon,
	)
	if err != nil {
		return err
	}

	fmt.Println(data.Name, "was caught!")
	CaughtPokemon[data.Name] = true

	return nil
}
