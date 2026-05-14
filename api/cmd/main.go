package main

import (
	"log"

	"github.com/farzadamr/go-clean-api/internal/env"
)

func main() {
	envCfg := &env.EnvFileConfig{
		EnvFile: ".env",
		//EnvFiles:     []string{".env.local", ".env"},
		RequiredVars: []string{"PORT"},
	}

	if err := env.Load(envCfg); err != nil {
		log.Printf("warning: cloud not load env files: %v", err)
	}

	cfg := &config{
		addr: ":" + env.GetEnv("PORT", "8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
