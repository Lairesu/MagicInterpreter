package main

import (
	"fmt"
	"os"
	"os/user"
	"MagicInterpreter/repl"
)

func main() {
	user, err:= user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, THIS IS MAGIC PROGRAMMING LANGUAGE!\n", user.Username)
	fmt.Printf("TYPE THE COMMANDS\n")
	repl.Start(os.Stdin,os.Stdout)
}