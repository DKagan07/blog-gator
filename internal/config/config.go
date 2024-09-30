package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFileName = "gatorconfig.json"

type Config struct {
	DbURL       string `json:"db_url,omitempty"`
	CurrentUser string `json:"current_user_name,omitempty"`
}

func ReadConfig() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	configPath := fmt.Sprintf("%s/.%s", home, ConfigFileName)

	cfg, err := getConfigFromPath(configPath)
	if err != nil {
		fmt.Printf("cannot find config file at path '%s'. Error: %v", configPath, err)
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUser = username

	configPath, err := getHomeDir()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0o744)
	if err != nil {
		return err
	}

	return nil
}

func getHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := fmt.Sprintf("%s/.%s", home, ConfigFileName)
	return configPath, nil
}

func getConfigFromPath(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("cannot find file at path: ", path)
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		fmt.Println("cannot unmarshal json")
		return Config{}, err
	}

	return cfg, nil
}
