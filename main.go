package main

import (
	"fmt"
	"os"

	"github.com/mnisyif/aggreGator/internal/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("couldnt read config: %s\n", err)
		os.Exit(1)
	}

	config.SetUser("Murtadha")

	fmt.Printf("url: %s\nuser: %s\n", config.DBURL, config.CurrentUser)
}
