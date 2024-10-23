package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type BattleRequest struct {
	AttackerID string `json:"attacker_id"`
	DefenderID string `json:"defender_id"`
}

func SubmitBattleHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var battleRequest BattleRequest
		if err := json.NewDecoder(r.Body).Decode(&battleRequest); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Marshal BattleRequest to JSON
		battleJSON, err := json.Marshal(battleRequest)
		if err != nil {
			http.Error(w, "Failed to serialize battle request", http.StatusInternalServerError)
			return
		}

		// Push battle request to Redis queue
		err = redisClient.RPush(ctx, "battleQueue", battleJSON).Err()
		if err != nil {
			http.Error(w, "Failed to submit battle", http.StatusInternalServerError)
			return
		}

		log.Printf("Battle request submitted: Attacker - %s, Defender - %s", battleRequest.AttackerID, battleRequest.DefenderID)
		w.WriteHeader(http.StatusAccepted)
	}
}
