package main

import (
	"fmt"
)

func CommandPokedex(...string) error {
	fmt.Println("Your Pokedex:")

	if len(CaughtPokemon) == 0 {
		fmt.Println("your Pokedex is empty!")
		return nil
	}

	for pokemon := range CaughtPokemon {
		fmt.Println(" -", pokemon)
	}

	return nil
}
