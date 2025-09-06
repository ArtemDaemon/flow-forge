package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Input struct {
		File string `yaml:"file"`
		Type string `yaml:"type"`
	} `yaml:"input"`
}

func ParseConfigFile(configPath string) (*Config, error) {
	var config Config

	log.Println("Start reading config file...")
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	log.Println("Start parsing YAML config...")
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	if config.Input.File == "" {
		return nil, fmt.Errorf("parsing config error: 'input.file' field is required")
	}
	if config.Input.Type == "" {
		return nil, fmt.Errorf("parsing config error: 'input.type' field is required")
	}

	return &config, nil
}
