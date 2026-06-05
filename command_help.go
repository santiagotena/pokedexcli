package main

import "fmt"

func commandHelp(cfg *config, param1 string) error {
	helpMessage :=
		"Welcome to the Pokedex!\n" +
			"Usage:\n\n" +
			"help: Displays a help message\n" +
			"map: Displays the next location areas\n" +
			"mapb: Displays the previous location areas\n" +
			"explore <area_name>: Displays the Pokemon available in the input and available area\n" +
			"catch <pokemon>: Attempts to catch a Pokemon\n" +
			"inspect <pokemon>: Details the stats of a caught Pokemon\n" +
			"pokedex: List all the caught Pokemon\n" +
			"exit: Exit the Pokedex\n"

	fmt.Print(helpMessage)
	return nil
}
