package main

import (
	"errors"
	"fmt"
	"os"
)

var localtions LocationAreas
var exploredLocation ExploredLocation

func commandHelp(args []string) error {
	commandMap := getCommandMap()
	if len(commandMap) == 0 {
		return errors.New("No commads available!")
	}
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func commandExit(args []string) error {
	os.Exit(0)
	return nil
}

func commandMap(args []string) error {
	err := assignLocationAreas(&localtions, "next")
	if err != nil {
		return err
	}
	printLocations()
	return nil
}

func commandMapb(args []string) error {
	err := assignLocationAreas(&localtions, "previous")
	if err != nil {
		return err
	}
	printLocations()
	return nil
}

func explore(args []string) error {
	// slog.Info("expore called with arg :", args[0])
	err := assignExploredLocation(&exploredLocation, args[0])
	if err != nil {
		return err
	}
	printExploredLocation()
	return nil
}

func getCommandMap() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 50 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 50 localtions",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explores choosen location",
			callback:    explore,
		},
	}
}

func printLocations() {
	for _, value := range localtions.Results {
		fmt.Println(value.Name)
	}
}

func printExploredLocation() {
	for _, value := range exploredLocation.PokemonEncounters {
		fmt.Println(value.Pokemon.Name)
	}
}
