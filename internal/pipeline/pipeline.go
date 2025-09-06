package pipeline

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Input struct {
		File string `yaml: "file"`
		Type string `yaml: "type"`
	} `yaml: "input"`
	Pipeline struct{} `yaml: "pipeline"`
	Output   struct{} `yaml: "output"`
}

func Run(configPath string) error {
	var config Config

	log.Println("Start reading config file...")
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return err
	}

	fmt.Println("🚀 Starting FlowForge pipeline...")
	fmt.Printf("📁 Using config: %s\n", configPath)
	return nil
}
