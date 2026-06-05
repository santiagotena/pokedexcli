package main

import (
	"fmt"
	"strconv"
)

func commandInspect(cfg *config, pokemon string) error {
	if pokemon == "" {
		return nil
	}

	if _, ok := cfg.pokedex[pokemon]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemonProfile := cfg.pokedex[pokemon]

	inspectMessage :=
		"Name: " + pokemon + "\n" +
			"Height: " + strconv.Itoa(pokemonProfile.height) + "\n" +
			"Weight: " + strconv.Itoa(pokemonProfile.weight) + "\n" +
			"Stats:\n" +
			"  -hp: " + strconv.Itoa(pokemonProfile.stats["hp"]) + "\n" +
			"  -attack: " + strconv.Itoa(pokemonProfile.stats["attack"]) + "\n" +
			"  -defense: " + strconv.Itoa(pokemonProfile.stats["defense"]) + "\n" +
			"  -special-attack: " + strconv.Itoa(pokemonProfile.stats["special-attack"]) + "\n" +
			"  -special-defense: " + strconv.Itoa(pokemonProfile.stats["special-defense"]) + "\n" +
			"  -speed: " + strconv.Itoa(pokemonProfile.stats["speed"]) + "\n" +
			"Types:\n"

	fmt.Print(inspectMessage)
	for _, pokemonType := range pokemonProfile.types {
		fmt.Print("  - " + pokemonType + "\n")
	}

	return nil
}
