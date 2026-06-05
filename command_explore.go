package main

import (
	"fmt"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		return nil
	}

	_, pokemons, ok := cfg.cache.Get(areaName)
	if ok {
		fmt.Println("Exploring " + areaName + "...")
		fmt.Println("Found Pokemon:")
		for _, pokemon := range pokemons {
			fmt.Println(" - " + pokemon)
		}
		return nil
	}

	url := pokeapi.BaseURL + pokeapi.LocationAreasEndpoint + areaName
	pokemons, err := pokeapi.FetchArea(url, areaName)
	if err != nil {
		return err
	}

	cfg.cache.Add(areaName, nil, pokemons)

	return nil
}
