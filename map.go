package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	LocationAreas []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
}

var page = -1

func CommandMapNext(...string) error {
	page++

	err := printMap(page)
	if err != nil {
		return err
	}

	return nil
}

func CommandMapBack(...string) error {
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
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%d&offset=%d", 20, page*20)

	if data, ok := GlobalCache.Get(url); ok {
		return processMapData(data)
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
	return processMapData(data)
}

func processMapData(data []byte) error {
	response := locationAreaResponse{}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	for _, area := range response.LocationAreas {
		fmt.Println(area.Name)
	}

	return nil
}
