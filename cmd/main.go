package main

import (
	"flag"
	"flow-forge/internal/pipeline"
	"log"
	"os"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to the config YAML file (required)")
	flag.StringVar(&configPath, "c", "", "Path to the config YAML file (required)")

	flag.Parse()

	// Checking for config file
	if configPath == "" {
		log.Println("Error: the required flag '--config' was not provided")
		flag.Usage()
		os.Exit(1)
	}

	if err := pipeline.Run(configPath); err != nil {
		log.Fatal(err)
	}
}
