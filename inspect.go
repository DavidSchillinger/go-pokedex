package main

import (
	"fmt"
)

type pokemonDetailResponse struct {
	Name   string        `json:"name"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Stats  []pokemonStat `json:"stats"`
	Types  []pokemonType `json:"types"`
}

type pokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	}
}

type pokemonType struct {
	Type struct {
		Name string `json:"name"`
	}
}

func CommandInspect(args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name of a Pokemon!")
		return nil
	}

	if _, ok := CaughtPokemon[args[0]]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	return printInspect(args[0])
}

func printInspect(pokemon string) error {
	data, err := Fetch[pokemonDetailResponse](
		"https://pokeapi.co/api/v2/pokemon/" + pokemon,
	)
	if err != nil {
		return err
	}

	fmt.Println("Name:", data.Name)
	fmt.Println("Height:", data.Height)
	fmt.Println("Weight:", data.Weight)
	fmt.Println("Stats:")
	for _, stat := range data.Stats {
		fmt.Println(" -", stat.Stat.Name+":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range data.Types {
		fmt.Println(" -", typ.Type.Name)
	}

	return nil
}
