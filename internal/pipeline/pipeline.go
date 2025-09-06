package pipeline

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

func Run(configPath string) error {
	var config Config

	log.Println("Start reading config file...")
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	log.Println("Start parsing YAML config...")
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return err
	}

	fmt.Println("🚀 Starting FlowForge pipeline...")
	return nil
}
