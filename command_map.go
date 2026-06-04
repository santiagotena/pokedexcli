package main

import (
	"encoding/json"
	"fmt"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := ""
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.nextLocationsURL
	}

	cache, ok := cfg.cache.Get(url)
	if ok {
		locationsArea := pokeapi.LocationArea{}
		err := json.Unmarshal(cache, &locationsArea)
		if err != nil {
			fmt.Println("Error marshalling response body")
			return err
		}
		for _, result := range locationsArea.Results {
			fmt.Println(result.Name)
		}
		cfg.prevLocationsURL = locationsArea.Previous
		cfg.nextLocationsURL = locationsArea.Next
		return nil
	}

	prev, next, dat, err := pokeapi.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = prev
	cfg.nextLocationsURL = next
	cfg.cache.Add(url, dat)

	return nil
}
