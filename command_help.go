package main

import "fmt"

func Help() error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println()
	fmt.Println("Commands:")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Println()

	return nil
}
