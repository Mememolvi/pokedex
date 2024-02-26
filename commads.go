package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

var localtions LocationAreas
var exploredLocation ExploredLocation

var caughtPokemonMap map[string]Pokemon = make(map[string]Pokemon)

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

func catch(args []string) error {
	// slog.Info("catch command called with ar :", args[0])
	pokemon, err := FetchPokemon(args[0])
	if err != nil {
		return err
	}
	attemptCatch(pokemon)
	return nil
}

func inspect(args []string) error {
	// slog.Info("inspect command called with ar :", args[0])
	pokemon, ok := caughtPokemonMap[args[0]]
	if ok {
		printPokemonDetails(&pokemon)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func printPokemonDetails(pokemon *Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf("-%v\n", v.Type.Name)
	}

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
		"catch": {
			name:        "catch",
			description: "attempt to catch choosen pokemon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "print extended pokemon details",
			callback:    inspect,
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

func attemptCatch(pokemon Pokemon) {
	fmt.Println("Throwing a Pokeball at pikachu...")
	exp := pokemon.BaseExperience
	rand := rand.Intn(exp)
	if rand < 100 {
		//caught
		caughtPokemonMap[pokemon.Name] = pokemon
		fmt.Println("pikachu was caught!")
	} else {
		//excaped
		fmt.Println("pikachu escaped!")
	}
}
