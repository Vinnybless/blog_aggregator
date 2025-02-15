package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}

	s.cfg.SetUser(cmd.args[0])

	fmt.Println(cmd.args[0], "has been set")

	return nil
}
