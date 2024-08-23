package main

import (
	"fmt"
	"os"
)

func (r *Repl) Quit() error {
	fmt.Println("See ya!")
	os.Exit(0)
	return nil
}
