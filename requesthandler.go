package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/Mememolvi/pokedexcli/internal/pokecache"
)

var c pokecache.Cache = pokecache.NewCache(time.Second * 10)

func assignLocationAreas(locations *LocationAreas, direction string) error {
	if direction == "previous" && locations.Previous == nil {
		return errors.New("Previous page doesnt exist!")
	}
	var url string
	if locations.Next == "" {
		// first fetch populate url
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=5"
	} else {
		if direction == "next" {
			url = localtions.Next
		} else {
			url = *locations.Previous
		}
	}

	body, err := fetchFromApi(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, locations)
	if err != nil {
		return err
	}
	return nil
}

func assignExploredLocation(exploredLocation *ExploredLocation, loactionName string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + loactionName
	body, err := fetchFromApi(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, exploredLocation)
	if err != nil {
		return err
	}
	return nil
}

func FetchPokemon(pokemonName string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pokemon := Pokemon{}
	body, err := fetchFromApi(url)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func fetchFromApi(url string) ([]byte, error) {
	if v, ok := c.Get(url); ok {
		return v, nil
	} else {
		res, err := http.Get(url)

		if err != nil {
			return nil, errors.New("FAILED TO MAKE API CALL")
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, errors.New("Response failed with status code: " + string(rune(res.StatusCode)) + " and\nbody:" + string(body) + "\n")
		}
		if err != nil {
			return nil, errors.New("FAILED TO MAKE API CALL")
		}
		go c.Add(url, body)
		return body, err
	}

}
