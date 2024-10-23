package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func GetLeaderboardHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		// Fetch leaderboard data from Redis (assuming sorted set)
		leaderboard, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
		if err != nil {
			http.Error(w, "Failed to retrieve leaderboard", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(leaderboard)
	}
}
