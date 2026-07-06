package main

import (
	"fmt"

	"github.com/mnisyif/aggreGator/internal/commands"
)

func handlerLogin(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a username argument")
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Username has been set successfully")

	return nil
}
