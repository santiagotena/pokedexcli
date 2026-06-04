package main

import (
	"encoding/json"
	"fmt"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.prevLocationsURL

	cache, ok := cfg.cache.Get(url)
	if ok {
		locationAreas := pokeapi.LocationAreas{}
		err := json.Unmarshal(cache, &locationAreas)
		if err != nil {
			fmt.Println("Error marshalling response body")
			return err
		}
		for _, result := range locationAreas.Results {
			fmt.Println(result.Name)
		}
		cfg.prevLocationsURL = locationAreas.Previous
		cfg.nextLocationsURL = locationAreas.Next
		return nil
	}

	prev, next, _, err := pokeapi.FetchLocationAreas(url)
	if err != nil {
		return err
	}
	cfg.prevLocationsURL = prev
	cfg.nextLocationsURL = next

	return nil
}
