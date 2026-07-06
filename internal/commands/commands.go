// Package commands is meant to allow users register cli commands and run them
package commands

import (
	"errors"

	"github.com/mnisyif/aggreGator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CliCommands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if s == nil {
		return errors.New("state is nil, cannot access config")
	}

	f, exists := c.CliCommands[cmd.Name]
	if !exists {
		return errors.New("command is not found")
	}

	return f(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CliCommands[name] = f
}
