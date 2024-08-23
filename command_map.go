package main

import (
	"errors"
	"fmt"

	"github.com/andrew-hayworth22/pokedex/pokedexapi"
)

var MAP_PAGE int

func (r *Repl) Map() error {
	MAP_PAGE++

	locations, err := r.client.GetLocationsPage(MAP_PAGE)
	if err != nil {
		return errors.New(err.Error())
	}

	printLocationsPage(locations)

	return nil
}

func (r *Repl) MapB() error {
	if MAP_PAGE < 2 {
		return errors.New("cannot print previous page")
	}

	MAP_PAGE--

	locations, err := r.client.GetLocationsPage(MAP_PAGE)
	if err != nil {
		return errors.New(err.Error())
	}

	printLocationsPage(locations)

	return nil
}

func printLocationsPage(locations []pokedexapi.Location) {
	fmt.Printf("--- Location Page %d ---\n", MAP_PAGE)
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
