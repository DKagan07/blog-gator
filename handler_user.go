package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Commands) == 0 {
		return fmt.Errorf("expect: %s <name>", cmd.Name)
	}

	newName := cmd.Commands[0]
	err := s.cfg.SetUser(newName)
	if err != nil {
		return fmt.Errorf("could not set user, error: %+v", err)
	}

	fmt.Printf("User %s has been set", newName)
	return nil
}
