package main

import (
	"fmt"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.prevLocationsURL

	prev, next, _, err := pokeapi.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = prev
	cfg.nextLocationsURL = next

	return nil
}
