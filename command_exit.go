package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, param1 string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
