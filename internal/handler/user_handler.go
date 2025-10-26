package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	logger := r.Context().Value("logger").(*slog.Logger)

	if len(username) == 0 {
		WriteJson(w, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := h.services.GetUserByUsername(r.Context(), username)

	if err != nil {
		logger.Error("failed to get user", "error", err)

		WriteErrorResponse(w, err)
		return
	}

	WriteJson(w, http.StatusOK, user)
}
