package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"

	"battles/pkg/models"
)

func CreatePlayerHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var player models.Player
		if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Validate player details (e.g., name length <= 20, etc.)
		if len(player.Name) > 20 {
			http.Error(w, "Name too long", http.StatusBadRequest)
			return
		}

		// Marshal player struct to JSON
		playerJSON, err := json.Marshal(player)
		if err != nil {
			http.Error(w, "Failed to serialize player", http.StatusInternalServerError)
			return
		}

		// Store player details in Redis
		playerKey := "player:" + player.ID
		err = redisClient.Set(ctx, playerKey, playerJSON, 0).Err()
		if err != nil {
			http.Error(w, "Failed to store player", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
