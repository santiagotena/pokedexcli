package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/santiagotena/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" {
		return nil
	}

	pokemonProfile := PokemonProfile{}
	completePokemonProfile := pokeapi.CompletePokemonProfile{}
	isPokemonFound := false
	err := error(nil)

	pokemon = strings.ToLower(pokemon)
	if _, ok := cfg.pokedex[pokemon]; ok {
		pokemonProfile = cfg.pokedex[pokemon]
	} else {
		url := pokeapi.BaseURL + pokeapi.PokemonEndpoint + pokemon
		completePokemonProfile, isPokemonFound, err = pokeapi.FetchPokemonProfile(url)
		if err != nil {
			return err
		}
		if !isPokemonFound {
			fmt.Println(pokemon + " is not a Pokemon")
			return nil
		}

		pokemonProfile.height = completePokemonProfile.Height
		pokemonProfile.weight = completePokemonProfile.Weight
		pokemonProfile.baseExperience = completePokemonProfile.BaseExperience
		for _, pokemonType := range completePokemonProfile.Types {
			pokemonProfile.types = append(pokemonProfile.types, pokemonType.Type.Name)
		}

		pokemonProfile.stats = make(map[string]int)
		stats := []int{}
		for _, pokemonStats := range completePokemonProfile.Stats {
			stats = append(stats, pokemonStats.BaseStat)
		}
		pokemonProfile.stats["hp"] = stats[0]
		pokemonProfile.stats["attack"] = stats[1]
		pokemonProfile.stats["defense"] = stats[2]
		pokemonProfile.stats["special-attack"] = stats[3]
		pokemonProfile.stats["special-defense"] = stats[4]
		pokemonProfile.stats["speed"] = stats[5]

		cfg.pokedex[pokemon] = pokemonProfile
	}

	fmt.Println("Throwing a Pokeball at " + pokemon + "...")

	isPokemonCaught := attemptToCatch(pokemonProfile.baseExperience)
	if isPokemonCaught {
		fmt.Println(pokemon + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Println(pokemon + " escaped!")
	return nil
}

func attemptToCatch(baseExperience int) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	baseCatchChance := r.Float64()
	catchPenalty := float64(baseExperience) / 1000.0

	if baseCatchChance-catchPenalty >= 0.5 {
		return true
	}
	return false
}
