package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocationAreas(url string) (*string, *string, []byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching location areas")
		return nil, nil, nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil, nil, nil, err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		fmt.Println("Error unmarshalling response body")
		return nil, nil, nil, err
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	prev := locationAreas.Previous
	next := locationAreas.Next

	return prev, next, dat, nil
}

func FetchArea(url string, areaName string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching areas")
		return nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil, err
	}

	area := Area{}
	err = json.Unmarshal(dat, &area)
	if err != nil {
		fmt.Println("Error unmarshalling response body")
		return nil, err
	}

	pokemons := []string{}
	fmt.Println("Exploring " + areaName + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range area.PokemonEncounters {
		pokemon := pokemonEncounter.Pokemon.Name
		fmt.Println(" - " + pokemon)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func FetchPokemonProfile(url string) (CompletePokemonProfile, bool, error) {
	completePokemonProfile := CompletePokemonProfile{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching areas")
		return CompletePokemonProfile{}, false, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return CompletePokemonProfile{}, false, nil
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return CompletePokemonProfile{}, false, err
	}

	err = json.Unmarshal(dat, &completePokemonProfile)
	if err != nil {
		fmt.Println("Error unmarshalling response body")
		return CompletePokemonProfile{}, false, err
	}

	return completePokemonProfile, true, nil
}
