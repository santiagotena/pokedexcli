package main

import "fmt"

func commandPokedex(cfg *config, param1 string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("No Pokemon has been caught yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.pokedex {
		fmt.Println(" - " + k)
	}
	return nil
}
