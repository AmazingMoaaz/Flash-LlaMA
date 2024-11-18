package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config structure to hold the Groq API key
type Config struct {
	GroqAPIKey string `json:"groq_api_key"` // Only Groq API Key
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(configFilePath string) (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}

	return &config, nil
}
