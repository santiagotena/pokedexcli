package main

import "fmt"

func commandHelp(cfg *config, param1 string) error {
	helpMessage :=
		"Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nmap: Displays the next location areas\nmapb: Displays the previous location areas\nexplore <area_name>: Displays the Pokemons available in the input and available area\nexit: Exit the Pokedex"

	fmt.Println(helpMessage)
	return nil
}
