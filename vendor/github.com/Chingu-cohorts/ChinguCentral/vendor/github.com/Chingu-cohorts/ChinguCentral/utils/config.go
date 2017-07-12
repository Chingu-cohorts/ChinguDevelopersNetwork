package utils

import (
	"encoding/json"
	"log"
	"os"
)

// Config represents the data contained in the config.json file
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
	Port string `json:"port"`
}

// LoadSettings loads the configuration for the application
func LoadSettings(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		log.Fatalf("Error reading the config file: %v", err)
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
