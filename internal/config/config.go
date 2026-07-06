// Package config that holds config for gator
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const confiFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
}

func Read(filename *string) (*Config, error) {
	if filename == nil {
		return nil, errors.New("please define a filename")
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Couldnt retrieve home address: %s\n", err)
		os.Exit(1)
	}

	configPath := fmt.Sprintf("%s/%s", homePath, *filename)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SetUser(user string) error {
	return nil
}
