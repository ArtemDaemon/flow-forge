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

func ParseConfigFile(path string) (*Config, error) {
	var c Config

	log.Println("Start reading config file...")
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.Println("Start parsing YAML config...")
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	if c.Input.File == "" {
		return nil, fmt.Errorf("parsing config error: 'input.file' field is required")
	}
	if c.Input.Type == "" {
		return nil, fmt.Errorf("parsing config error: 'input.type' field is required")
	}
	if c.Input.Type != "csv" {
		return nil, fmt.Errorf("parsing config error: 'input.type' field must have value 'csv'")
	}

	return &c, nil
}
