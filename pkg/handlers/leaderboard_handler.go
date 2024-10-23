package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"battles/pkg/services"
)

type LeaderboardHandler struct {
	gameService *services.GameService
}

func NewLeaderboardHandler(service *services.GameService) *LeaderboardHandler {
	return &LeaderboardHandler{
		gameService: service,
	}
}

func (h *LeaderboardHandler) GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	leaderboard, err := h.gameService.GetLeaderboard(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve leaderboard", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(leaderboard)
}
