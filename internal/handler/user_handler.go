package handler

import (
	"context"
	"encoding/json"
	"golang-error/internal/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("id")

	user, err := h.service.GetByUsername(context.Background(), username)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
