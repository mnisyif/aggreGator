package main

import (
	"fmt"
	"os"

	"github.com/mnisyif/aggreGator/internal/config"
)

func main() {
	filename := ".gatorconfig.json"

	config, err := config.Read(&filename)
	if err != nil {
		fmt.Printf("couldnt read config: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("DB URL: %s\n", config.DBURL)
}
