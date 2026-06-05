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
	callback    func(*config, string) error
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
}

func replLoop() {
	commands := getCommands()
	config := &config{cache: pokecache.NewCache(5 * time.Minute)}
	runRepl(commands, config)
}

func getCommands() map[string]cliCommand {
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
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemons in the selected area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Details a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandPokedex,
		},
	}

	return commands
}

func runRepl(commands map[string]cliCommand, config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		param1 := ""
		if len(cleanedInput) >= 2 {
			param1 = cleanedInput[1]
		}

		if _, ok := commands[command]; ok {
			err := commands[command].callback(config, param1)
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
