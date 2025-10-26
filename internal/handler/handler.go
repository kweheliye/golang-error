package handler

import (
	"golang-error/internal/middlewear"
	"golang-error/internal/service"
	"log/slog"
	"net/http"
)

type Handler struct {
	services service.Service
	mux      *http.ServeMux
	logger   *slog.Logger
}

func NewHandler(services *service.Service, logger *slog.Logger) *Handler {
	h := &Handler{
		services: *services,
		mux:      http.NewServeMux(),
		logger:   logger,
	}

	h.registerRoutes()
	return h
}

func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/users/", middlewear.AttachLogger(h.GetUserByUsername, h.logger))
}

func (h *Handler) Router() http.Handler {
	return h.mux
}
