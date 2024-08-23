package main

import "fmt"

func (r *Repl) Help() error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println()
	fmt.Println("Commands:")

	for _, command := range r.GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Println()

	return nil
}
