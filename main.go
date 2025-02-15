package main

import (
	"fmt"
	"os"

	"github.com/Vinnybless/blog_aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	s := state{&cfg}

	c := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	c.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	if err := c.run(&s, command{
		name: args[1],
		args: args[2:],
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
