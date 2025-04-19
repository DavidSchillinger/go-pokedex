package main

import (
	"fmt"
)

type locationAreaResponse struct {
	LocationAreas []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
}

var page = -1

func CommandMapNext(_ ...string) error {
	page++

	err := printMap(page)
	if err != nil {
		return err
	}

	return nil
}

func CommandMapBack(_ ...string) error {
	if page == 0 {
		fmt.Println("You're on the first page.")
		return nil
	}

	page--

	err := printMap(page)
	if err != nil {
		return err
	}

	return nil
}

func printMap(page int) error {
	data, err := Fetch[locationAreaResponse](
		fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%d&offset=%d", 20, page*20),
	)
	if err != nil {
		return err
	}

	for _, area := range data.LocationAreas {
		fmt.Println(area.Name)
	}

	return nil
}
