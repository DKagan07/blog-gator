package main

import (
	"errors"
)

type command struct {
	Name     string
	Commands []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	v, ok := c.cmds[cmd.Name]
	if !ok {
		return errors.New("command not registered")
	}
	return v(s, cmd)
}
