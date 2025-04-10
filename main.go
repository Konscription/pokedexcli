package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	splitFn := func(c rune) bool { return c == ' ' }
	output := strings.FieldsFunc(strings.ToLower(text), splitFn)
	for i := range len(output) - 1 {
		output[i] = strings.TrimSpace(output[i])
	}
	return output
}

func main() {
	//scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		//wait for user input (enter key pressed)
		if !scanner.Scan() {
			// if scan fails exit loop
			fmt.Println("\nExiting Pokedex...")
			break
		}

		// get and clean input text
		input := cleanInput(scanner.Text())

		// skip if empty
		if len(input) == 0 {
			continue
		}

		// split and get first word
		command := input[0]

		// print command
		fmt.Printf("Your command was: %s\n", command)
	}
}
