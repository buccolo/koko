package main

import (
	"fmt"
	"koko/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\nThis is Koko, an interpreter for the Monkey language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
