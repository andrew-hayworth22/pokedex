package main

import (
	"errors"
	"fmt"

	"github.com/andrew-hayworth22/pokedex/pokedexapi"
)

func (r *Repl) Explore() error {
	location := r.arguments[0]
	fmt.Printf("Exploring %s...\n", location)

	pokemon, err := r.client.GetPokemonByLocation(location)
	if err != nil {
		return errors.New(err.Error())
	}

	printPokemonList(pokemon)

	return nil
}

func printPokemonList(pokemon []pokedexapi.Pokemon) {
	for _, poke := range pokemon {
		fmt.Printf(" - %s\n", poke.Name)
	}
}
