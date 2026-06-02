package main

import (
	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := ""
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.nextLocationsURL
	}

	prev, next, err := pokeapi.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = prev
	cfg.nextLocationsURL = next

	return nil
}
