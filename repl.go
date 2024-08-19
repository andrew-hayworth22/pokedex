package main

import (
	"fmt"
)

type Command struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays instructions for using the Pokedex",
			Callback:    Help,
		},
		"quit": {
			Name:        "quit",
			Description: "Quits the application",
			Callback:    Quit,
		},
	}
}

func HandleCommand(command string) error {
	commands := GetCommands()

	cmd, ok := commands[command]
	if !ok {
		return fmt.Errorf("invalid command: %s", command)
	}

	return cmd.Callback()
}
