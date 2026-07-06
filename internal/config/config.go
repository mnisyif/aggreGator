// Package config that holds config for gator
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func GetConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
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
	if user == "" {
		return fmt.Errorf("enter a name to update the config")
	}
	c.CurrentUser = user

	err := write(c)
	if err != nil {
		return err
	}

	return err
}

func write(cfg *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, prettyJSON, 0o600)

	return err
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("couldnt retrieve home address: %s", err)
	}

	return fmt.Sprintf("%s/%s", homePath, configFileName), nil
}
