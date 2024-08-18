package main

import (
	"fmt"
	"os"

	"github.com/AlvaroPrates/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

type config struct {
	next     string
	previous *string
}

func NewConfig() *config {
	return &config{}
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Shows locations from the pokemon world",
			Callback:    commandLocationAreas,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println()

	return nil
}

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandLocationAreas(cfg *config) error {
	payload, err := pokeapi.GetLocationAreas()
	if err != nil {
		fmt.Printf("Failed to get location areas: %s\n", err.Error())
	}

	cfg.next = payload.Next
	cfg.previous = payload.Previous

	for _, locationArea := range payload.LocationAreas {
		fmt.Println(locationArea.Location)
	}

	return nil
}
