package main

import (
	"context"

	"github.com/mnisyif/aggreGator/internal/commands"
	"github.com/mnisyif/aggreGator/internal/database"
)

func middlewareLoggedIn(handler func(s *commands.State, cmd commands.Command, user database.User) error) func(*commands.State, commands.Command) error {
	return func(s *commands.State, c commands.Command) error {
		user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
		if err != nil {
			return err
		}

		return handler(s, c, user)
	}
}
