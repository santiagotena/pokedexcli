package main

import (
	"encoding/json"
	"fmt"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, param1 string) error {
	url := ""
	if cfg.nextLocationsURL == nil {
		url = pokeapi.BaseURL + pokeapi.LocationAreasEndpoint
	} else {
		url = *cfg.nextLocationsURL
	}

	cache, _, ok := cfg.cache.Get(url)
	if ok {
		locationAreas := pokeapi.LocationAreas{}
		err := json.Unmarshal(cache, &locationAreas)
		if err != nil {
			fmt.Println("Error unmarshalling response body")
			return err
		}
		for _, result := range locationAreas.Results {
			fmt.Println(result.Name)
		}
		cfg.prevLocationsURL = locationAreas.Previous
		cfg.nextLocationsURL = locationAreas.Next
		return nil
	}

	prev, next, dat, err := pokeapi.FetchLocationAreas(url)
	if err != nil {
		return err
	}
	cfg.prevLocationsURL = prev
	cfg.nextLocationsURL = next
	cfg.cache.Add(url, dat, nil)

	return nil
}
