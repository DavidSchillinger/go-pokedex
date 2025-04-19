package main

import (
	"encoding/json"
	"fmt"
)

type pokemonDetailResponse struct {
	Name   string        `json:"name"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Stats  []pokemonStat `json:"stats"`
	Type   []pokemonType `json:"types"`
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

	err := printInspect(args[0])
	if err != nil {
		return err
	}

	return nil
}

func printInspect(pokemon string) error {
	data, err := CachedFetch("https://pokeapi.co/api/v2/pokemon/" + pokemon)
	if err != nil {
		return err
	}

	return processPokemonDetailData(data)
}

func processPokemonDetailData(data []byte) error {
	response := pokemonDetailResponse{}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	fmt.Println("Name:", response.Name)
	fmt.Println("Height:", response.Height)
	fmt.Println("Weight:", response.Weight)
	fmt.Println("Stats:")
	for _, stat := range response.Stats {
		fmt.Println(" -", stat.Stat.Name+":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range response.Type {
		fmt.Println(" -", typ.Type.Name)
	}

	return nil
}
