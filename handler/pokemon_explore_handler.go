package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/GabrielPereira187/pokedexcli/types"
)

const (
	base_url_explore = "https://pokeapi.co/api/v2/location-area/%s"
)


func GetExploreAreaPokemons(args ...string) error {
	name := args[1]
	resp, err := http.Get(fmt.Sprintf(base_url_explore, name))
	if err != nil {
		return errors.New("Error to get data")
	}
	defer resp.Body.Close()
	var exploreResponse types.ExploreResponse

	if err = json.NewDecoder(resp.Body).Decode(&exploreResponse); err != nil {
		return errors.New("Error to unmarshall")
	}

	log.Printf("Exploring %s\n", name)
	log.Println("Found Pokemon:")

	for _, pokemon := range exploreResponse.PokemonEncounters {
		log.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil

}
