package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ",", "")
	return strings.Fields(text)
}

func commandHelp() error {
	helpMessage :=
		"Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex"

	fmt.Println(helpMessage)
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("Error fetching locations")
		return err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return err
	}

	locationsArea := LocationArea{}
	err = json.Unmarshal(dat, &locationsArea)
	if err != nil {
		fmt.Println("Error marshalling response body")
		return err
	}

	for _, result := range locationsArea.Results {
		fmt.Println(result["name"])
	}

	return nil
}

func commandMapb() error {

	return nil
}

func replLoop() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays an instructions message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays the names of the previous 20 locations in the Pokemon world.",
			callback:    commandMapb,
		},
	}

	config := &config{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)

		command := cleanedInput[0]

		if _, ok := commands[command]; ok {
			err := commands[command].callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
