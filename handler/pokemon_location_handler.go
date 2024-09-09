package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GabrielPereira187/pokedexcli/types"
)

const (
	base_url_location = "https://pokeapi.co/api/v2/location-area/?offset=%v&limit=%v"
	qty = 20
)

var offset = 0

func GetNextLocationPage(args ...string) (error){
	if err := getLocations(offset); err != nil {
		return err
	}
	offset += qty
	return nil
}

func GetPreviousLocationPage(args ...string) (error){
	if offset == 0 {
		return errors.New("No more offset values")
	}
	offset -= qty
	if err := getLocations(offset); err != nil {
		return err
	}
	return nil
}


func getLocations(offset int) error {
	response , err := http.Get(fmt.Sprintf(base_url_location, offset, qty))
	if err != nil {
		return errors.New("Error to get locations")
	}
	defer response.Body.Close()

	var locations types.LocationResponse

	err = json.NewDecoder(response.Body).Decode(&locations)
	if err != nil {
		return errors.New("Error to unmarshal")
	}

	offset += qty

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}