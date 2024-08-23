package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrew-hayworth22/pokedex/pokedexapi"
)

type Command struct {
	Name        string
	Description string
	Callback    func() error
	Arguments   int
}

type Repl struct {
	client    pokedexapi.Client
	arguments []string
	pokedex   map[string]pokedexapi.Pokemon
}

func NewRepl() Repl {
	return Repl{
		client:  pokedexapi.NewClient(30*time.Minute, 30*time.Second),
		pokedex: map[string]pokedexapi.Pokemon{},
	}
}

func (r *Repl) GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays instructions for using the Pokedex",
			Callback:    r.Help,
			Arguments:   0,
		},
		"quit": {
			Name:        "quit",
			Description: "Quits the application",
			Callback:    r.Quit,
			Arguments:   0,
		},
		"map": {
			Name:        "map",
			Description: "Shows the next 20 locations",
			Callback:    r.Map,
			Arguments:   0,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous 20 locations",
			Callback:    r.MapB,
			Arguments:   0,
		},
		"explore": {
			Name:        "explore [location]",
			Description: "List the Pokemon in a location",
			Callback:    r.Explore,
			Arguments:   1,
		},
		"catch": {
			Name:        "catch [pokemon]",
			Description: "Try to catch a Pokemon",
			Callback:    r.Catch,
			Arguments:   1,
		},
		"inspect": {
			Name:        "inpect [pokemon]",
			Description: "Display the details of a Pokemon that you have captured",
			Callback:    r.Inspect,
			Arguments:   1,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all of the Pokemon that you have caught",
			Callback:    r.Pokedex,
			Arguments:   0,
		},
	}
}

func (r *Repl) HandleCommand(args ...string) error {
	if len(args) == 0 {
		return nil
	}
	commands := r.GetCommands()
	command := args[0]

	cmd, ok := commands[command]
	if !ok {
		return fmt.Errorf("invalid command: %s", command)
	}

	args = args[1:]
	if len(args) != cmd.Arguments {
		return fmt.Errorf("invalid arguments: given=%d need=%d", len(args), cmd.Arguments)
	}
	r.arguments = args

	return cmd.Callback()
}

func (r *Repl) StartRepl() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			return
		}

		args := cleanArgs(input)

		if err := r.HandleCommand(args...); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func cleanArgs(input []byte) []string {
	inputString := strings.ToLower(string(input))
	args := strings.Split(inputString, " ")
	return args
}
