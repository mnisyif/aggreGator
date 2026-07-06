package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mnisyif/aggreGator/internal/commands"
	"github.com/mnisyif/aggreGator/internal/database"
)

func handlerLogin(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a username argument\n")
	}

	user, err := s.DB.GetUserByName(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("username not found, please register to login\n")
	}
	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("Username has been set successfully")

	return nil
}

func handlerRegister(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("register command expects a username argument\n")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}

	user, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("could not register <%s> : %s\n", cmd.Args[0], err)
	}

	s.Cfg.SetUser(user.Name)
	fmt.Println("User was created successfully")
	fmt.Printf("ID: %s, Name: %s, Created At: %s\n", user.ID, user.Name, user.CreatedAt)
	return nil
}
