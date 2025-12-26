package main

import (
	"fmt"
	"github.com/lancelote/writing-a-compiler-in-go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! this is the Monkey programming language!\n", user.Username)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
