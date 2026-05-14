package handler

import (
	"net/http"
)

type HealthHandler struct {
	Message string
}

func NewHealthHandler(message string) *HealthHandler {
	return &HealthHandler{Message: message}
}

func (h *HealthHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Message))
}
