package main

import "fmt"

func (r *Repl) Pokedex() error {
	fmt.Println("Your Pokedex:")
	for key, _ := range r.pokedex {
		fmt.Printf("    - %s\n", key)
	}
	return nil
}
