package config

import (
	"encoding/json"
	"log"
	"os"
)

func getConfigPath() string {
	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path + "/.gatorconfig.json"
}

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(n string) {
	c.CurrentUserName = n

	path := getConfigPath()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	file.Write(jsonData)
}

func Read() Config {
	path := getConfigPath()

	file, err := os.OpenFile(path, os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cfg Config

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
