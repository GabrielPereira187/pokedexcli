package initializers

import (
	"fmt"

	"github.com/GabrielPereira187/pokedexcli/handler"
	"github.com/GabrielPereira187/pokedexcli/types"
)

func commandExit(args ...string) error {
	return nil
}

func commandHelp(args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func clearScreen(args ...string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

func LoadCommands() map[string]types.CliCommand{
	return map[string]types.CliCommand{
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
		"clear": {
			Name: "clear",
			Description: "Clear the console",
			Callback: clearScreen,
		},
		"map":{
			Name: "map",
			Description: "Display a page of names for location of Pokemons",
			Callback: handler.GetNextLocationPage,
		},
		"mapb":{
			Name: "mapb",
			Description: "Display the previous page of names for location of Pokemons",
			Callback: handler.GetPreviousLocationPage,
		},
		"explore":{
			Name: "explore",
			Description: "Expĺore pokemons of determined area",
			Callback: handler.GetExploreAreaPokemons,
		},
		"catch":{
			Name: "catch",
			Description: "Catch Pokemons",
			Callback: handler.CatchPokemon,
		},
		"inspect":{
			Name: "inspect",
			Description: "Obtain the description of the pokemons",
			Callback: handler.InspectPokemon,
		},
		"pokedex":{
			Name: "pokedex",
			Description: "Obtain the pokedex",
			Callback: handler.GetPokedex,
		},
	}
}
