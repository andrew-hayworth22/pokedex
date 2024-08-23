package main

import (
	"fmt"

	"github.com/andrew-hayworth22/pokedex/pokedexapi"
)

func (r *Repl) Inspect() error {
	pokemonName := r.arguments[0]
	pokemon, ok := r.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not captured a %s", pokemonName)
	}

	printPokemon(pokemon)

	return nil
}

func printPokemon(pokemon pokedexapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
}
