package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, exists := c.registeredCommands[cmd.Name]; exists {
		return f(s, cmd)
	}
	return fmt.Errorf("the command %s wasn't registered", cmd.Name)
}
