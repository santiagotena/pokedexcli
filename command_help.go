package main

import "fmt"

func commandHelp(cfg *config) error {
	helpMessage :=
		"Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex"

	fmt.Println(helpMessage)
	return nil
}
