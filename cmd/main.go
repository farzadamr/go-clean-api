package main

import (
	"log"

	"github.com/farzadamr/go-clean-api/api"
	"github.com/farzadamr/go-clean-api/config"
	"github.com/farzadamr/go-clean-api/pkg/logging"
)

func main() {
	// load config
	cfg := LoadAndParseConfig()
	if err := logging.Init(cfg); err != nil {
		log.Fatalf("logger can not initialized: %w", err)
	}
	logging.Debug("Starting The Server on:", "port", cfg.HTTP.Port)
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
