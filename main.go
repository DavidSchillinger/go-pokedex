package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args ...string) error
}

var commands map[string]cliCommand

func main() {
	commands = getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println(commands["exit"].name)
			break
		}

		input := cleanInput(scanner.Text())
		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		fmt.Println("")
		err := command.callback(input[1:]...)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			callback:    commandExit,
			description: "Exit the Pokedex",
		},
		"help": {
			name:        "help",
			callback:    commandHelp,
			description: "Displays a help message",
		},
		"map": {
			name:        "map",
			callback:    CommandMapNext,
			description: "Displays the names of the next 20 location areas",
		},
		"mapb": {
			name:        "mapb",
			callback:    CommandMapBack,
			description: "Displays the names of the previous 20 location areas",
		},
		"explore": {
			name:        "explore",
			callback:    CommandExplore,
			description: "Displays the names of Pokemon in the area",
		},
		"catch": {
			name:        "catch",
			callback:    CommandCatch,
			description: "Attempts to catch a Pokemon by name",
		},
		"inspect": {
			name:        "inspect",
			callback:    CommandInspect,
			description: "Inspect a Pokemon in your Pokedex",
		},
		"pokedex": {
			name:        "pokedex",
			callback:    CommandPokedex,
			description: "Shows a list of caught Pokemon",
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func commandExit(...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}

	return nil
}
