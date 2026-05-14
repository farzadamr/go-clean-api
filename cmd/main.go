package main

import (
	"log"

	"github.com/farzadamr/go-clean-api/api"
	"github.com/farzadamr/go-clean-api/config"
)

func main() {
	cfg := LoadAndParseConfig()

	log.Fatal(api.Run(cfg))
}

func LoadAndParseConfig() *config.Config {
	loader := config.NewLoader()
	if err := loader.LoadEnv(); err != nil {
		log.Fatalf("failed to load env: %v", err)
		return nil
	}

	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("invalid config: %v", err)
		return nil
	}

	return cfg
}
