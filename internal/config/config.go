package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	config_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	config_file, err := os.Open(config_path)
	if err != nil {
		return Config{}, err
	}
	defer config_file.Close()

	body, err := io.ReadAll(config_file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(user_name string) error {
	cfg.CurrentUserName = user_name
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home_dir, configFileName), nil
}

func write(cfg Config) error {
	config_path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	config_file, err := os.Create(config_path)
	if err != nil {
		return err
	}
	defer config_file.Close()

	encoder := json.NewEncoder(config_file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
