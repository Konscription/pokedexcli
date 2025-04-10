package main

import "strings"

func cleanInput(text string) []string {
	splitFn := func(c rune) bool { return c == ' ' }
	output := strings.FieldsFunc(strings.ToLower(text), splitFn)
	for i := range len(output) - 1 {
		output[i] = strings.TrimSpace(output[i])
	}
	return output
}

func main() {

}
