package router

import (
	"net/http"

	"github.com/farzadamr/go-clean-api/api/handler"
	"github.com/go-chi/chi/v5"
)

func healthRoutes() http.Handler {
	r := chi.NewRouter()
	h := handler.NewHealthHandler("Hello form clean-api")
	r.Get("/", h.GetMessage)

	return r
}
