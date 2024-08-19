package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			return
		}

		command := strings.ToLower(string(input))
		if err := HandleCommand(command); err != nil {
			fmt.Println(err.Error())
		}
	}
}
