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

		// check if command exists
		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cliCommand represents a command in the REPL
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapback,
		},
	}
}
