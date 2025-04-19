package main

import (
	"fmt"
)

func CommandPokedex(...string) error {
	err := printPokedex()
	if err != nil {
		return err
	}

	return nil
}

func printPokedex() error {
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
