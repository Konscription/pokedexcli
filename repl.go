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

func startRepl() {
	//scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// get and clean input text
		input := cleanInput(scanner.Text())

		// skip if empty
		if len(input) == 0 {
			continue
		}

		// get first word
		command := input[0]

		// print command
		fmt.Printf("Your command was: %s\n", command)
	}
}
