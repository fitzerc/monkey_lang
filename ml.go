package main

import (
	"fmt"
	"monkey_lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s. This is the Monkey programming language\n", user.Username)
	fmt.Println("Try Some Commands")

	repl.Start(os.Stdin, os.Stdout)
}
