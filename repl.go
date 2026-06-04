package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/santiagotena/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
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
	config.cache = pokecache.NewCache(5 * time.Minute)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)

		command := cleanedInput[0]

		if _, ok := commands[command]; ok {
			err := commands[command].callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ",", "")
	return strings.Fields(text)
}
