package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mnisyif/aggreGator/internal/commands"
	"github.com/mnisyif/aggreGator/internal/database"
	"github.com/mnisyif/aggreGator/internal/rss"
)

func handlerLogin(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a username argument  ")
	}

	user, err := s.DB.GetUserByName(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("username not found, please register to login  ")
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
		return fmt.Errorf("register command expects a username argument  ")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}

	user, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("could not register <%s> : %s  ", cmd.Args[0], err)
	}

	s.Cfg.SetUser(user.Name)
	fmt.Println("User was created successfully")
	fmt.Printf("ID: %s, Name: %s, Created At: %s\n", user.ID, user.Name, user.CreatedAt)
	return nil
}

func handlerReset(s *commands.State, cmd commands.Command) error {
	err := s.DB.ResetTable(context.Background())
	return err
}

func handlerUsers(s *commands.State, cmd commands.Command) error {
	usersList, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range usersList {
		if user.Name == s.Cfg.CurrentUser {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}

func handlerFeed(s *commands.State, cmd commands.Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	// if len(cmd.Args) == w {
	// 	return fmt.Errorf("agg command expects a link to fetch RSS feed from")
	// }

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}

func handlerAddFeed(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("addFeed expects <name_of_feed> and <url_of_feed> as arguments")
	}

	user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	fmt.Printf("Feed ID: %s\n", feed.ID)
	fmt.Printf("Feed Title: %s\n", feed.Name)
	fmt.Printf("Feed URL: %s\n", feed.Url)

	return nil
}

func handlerFeeds(s *commands.State, cmd commands.Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		user, err := s.DB.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("Feed creator: %s\n", user.Name)
		fmt.Println()
	}

	return nil
}
