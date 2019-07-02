package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/nicholasgriffen/es3-repl/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to ECMAScript 3. Type.\n%s!", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
