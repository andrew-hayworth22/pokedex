package main

import (
	"fmt"
	"math/rand"
)

var CATCH_THRESHOLD = 75

func (r *Repl) Catch() error {
	pokemonName := r.arguments[0]
	if _, ok := r.pokedex[pokemonName]; ok {
		fmt.Printf("You have already captured %s", pokemonName)
	}

	pokemon, err := r.client.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	randomInt := rand.Intn(pokemon.BaseExperience)
	if randomInt < CATCH_THRESHOLD {
		fmt.Printf("NO! %s got away!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("You caught %s!\n", pokemon.Name)
	r.pokedex[pokemon.Name] = pokemon

	return nil
}
