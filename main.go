package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/mnisyif/aggreGator/internal/commands"
	"github.com/mnisyif/aggreGator/internal/config"
	"github.com/mnisyif/aggreGator/internal/database"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("couldnt read config: %s\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		fmt.Printf("could not create a connection with database: %s\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	userState := commands.State{
		Cfg: config,
		DB:  dbQueries,
	}

	cmdList := commands.Commands{
		CliCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmdList.Register("login", handlerLogin)
	cmdList.Register("register", handlerRegister)
	cmdList.Register("reset", handlerReset)

	if len(os.Args) < 2 {
		fmt.Printf("You are missing command arguments\n")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmdList.Run(&userState, cmd)
	if err != nil {
		fmt.Printf("failed to run command <%s>: %s", cmd.Name, err)
		os.Exit(1)
	}
}
