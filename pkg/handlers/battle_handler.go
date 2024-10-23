package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"battles/pkg/models"
	"battles/pkg/services"
)

type BattleHandler struct {
	gameService *services.GameService
}

func NewBattleHandler(service *services.GameService) *BattleHandler {
	return &BattleHandler{
		gameService: service,
	}
}

func (h *BattleHandler) SubmitBattleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var battleRequest models.BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&battleRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.gameService.SubmitBattle(ctx, battleRequest); err != nil {
		http.Error(w, "Failed to submit battle request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
