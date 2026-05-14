package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MountRoutes() http.Handler {
	r := chi.NewRouter()

	r.Mount("/health", healthRoutes())

	return r
}
