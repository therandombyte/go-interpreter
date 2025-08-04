package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/therandombyte/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is nik!\n", user.Username)
	fmt.Print("Feel free to ask anything\n")
	repl.Start(os.Stdin, os.Stdout)
}
