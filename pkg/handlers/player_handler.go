package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"battles/pkg/models"
	"battles/pkg/services"
)

type PlayerHandler struct {
	gameService *services.GameService
}

func NewPlayerHandler(service *services.GameService) *PlayerHandler {
	return &PlayerHandler{
		gameService: service,
	}
}

func (h *PlayerHandler) CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var player models.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.gameService.CreatePlayer(ctx, player); err != nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
