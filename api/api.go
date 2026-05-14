package api

import (
	"net/http"

	"github.com/farzadamr/go-clean-api/api/router"
	"github.com/farzadamr/go-clean-api/config"
)

func Run(cfg *config.Config) error {
	handler := router.MountRoutes()

	srv := &http.Server{
		Addr:         ":" + cfg.HTTP.Port,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
		Handler:      handler,
	}

	return srv.ListenAndServe()
}
