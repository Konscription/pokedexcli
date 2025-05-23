package main

import (
	"fmt"
	"os"
)

// commandExit handles the 'exit' command
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // Unreachable, but satisfies function signature
}
