package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/GabrielPereira187/pokedexcli/types"
)

const (
	base_url_pokemon = "https://pokeapi.co/api/v2/pokemon/%s"
	chance_to_catch = 10
)


var pokemons = make(map[string]types.PokemonResponse)

func CatchPokemon(args ...string) error {
	chanceToCatch := rand.Intn(100)
	pokemonName := args[1]

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)

	if chanceToCatch > chance_to_catch {
		pokemon, err := getPokemonData(pokemonName);
		if  err != nil {
			return err
		}

		pokemons[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil

}

func getPokemonData(pokemonName string) (types.PokemonResponse, error) {
	var pokemonResponse types.PokemonResponse
	response, err := http.Get(fmt.Sprintf(base_url_pokemon, pokemonName))
	if err != nil {
		return types.PokemonResponse{}, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&pokemonResponse); err != nil {
		return types.PokemonResponse{}, err
	}

	return pokemonResponse, nil
}

func InspectPokemon(args ...string) error {
	pokemonName := args[1]

	if pokemon, ok := pokemons[pokemonName]; ok == false {
		fmt.Println("you have not caught that pokemon")
	} else {
		printPokemonData(pokemon)
	}

	return nil
}

func printPokemonData(pokemon types.PokemonResponse) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d \n", pokemon.Height)
	fmt.Printf("Weight: %d \n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("	%s\n", typ.Type.Name)
	}
}

func GetPokedex(args ...string) error {
	for key := range pokemons {
		fmt.Printf("- %s\n", key)
	} 

	return nil
}