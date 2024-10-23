package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"battles/database"
	"battles/pkg/handlers"
	"battles/pkg/services"
)

func main() {
	ctx := context.Background()
	// Initialize Redis client
	redisClient := database.NewRedisClient(ctx)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register endpoints
	r.HandleFunc("/player", handlers.CreatePlayerHandler(redisClient)).Methods("POST")
	r.HandleFunc("/battle", handlers.SubmitBattleHandler(redisClient)).Methods("POST")
	r.HandleFunc("/leaderboard", handlers.GetLeaderboardHandler(redisClient)).Methods("GET")

	// Run battle processor in a separate Goroutine
	go services.ProcessBattlesConcurrently(redisClient)

	// Start the server
	http.ListenAndServe(":8080", r)
}
