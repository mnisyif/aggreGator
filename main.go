package main

import (
	"fmt"
	"os"

	"github.com/mnisyif/aggreGator/internal/commands"
	"github.com/mnisyif/aggreGator/internal/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("couldnt read config: %s\n", err)
		os.Exit(1)
	}

	userState := commands.State{
		Cfg: config,
	}

	cmdList := commands.Commands{
		CliCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmdList.Register("login", handlerLogin)

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
