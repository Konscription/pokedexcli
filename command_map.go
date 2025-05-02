package main

import (
	"fmt"
)

// commandMap displays the names of 20 location areas
// in the pokemon world. Each subsequent call to map
// should display the next 20 location areas.
var currentPage int

func commandMap() error {
	locationAreas := getLocationAreas(currentPage)
	for _, area := range locationAreas[:] {
		fmt.Printf("%s\n", area)
	}
	currentPage++
	return nil
}

func commandMapback() error {
	if currentPage == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	currentPage -= 2
	if currentPage < 0 {
		currentPage = 0
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreas := getLocationAreas(currentPage)
	for _, area := range locationAreas[:] {
		fmt.Printf("%s\n", area)
	}
	return nil
}
