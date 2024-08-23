package main

import "fmt"

func (r *Repl) Pokedex() error {
	fmt.Println("Your Pokedex:")
	for key := range r.pokedex {
		fmt.Printf("    - %s\n", key)
	}
	return nil
}
