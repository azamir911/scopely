package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"battles/database"
	"battles/pkg/auth"
	"battles/pkg/handlers"
	"battles/pkg/repository"
	"battles/pkg/services"
)

func main() {
	ctx := context.Background()
	// Initialize Redis client
	redisClient := database.NewRedisClient(ctx)

	// Clear all databases in Redis
	//err := redisClient.FlushAll(ctx).Err()
	//if err != nil {
	//	panic(err)
	//}

	// Initialize repository, services, and handlers
	redisRepository := repository.NewRedisRepository(redisClient)
	gameService := services.NewGameService(redisRepository)
	battleService := services.NewBattleService(redisRepository)

	playerHandler := handlers.NewPlayerHandler(gameService)
	battleHandler := handlers.NewBattleHandler(gameService)
	leaderboardHandler := handlers.NewLeaderboardHandler(gameService)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register endpoints with authentication middleware
	r.Handle("/player", auth.AuthMiddleware(http.HandlerFunc(playerHandler.CreatePlayerHandler))).Methods("POST")
	r.Handle("/battle", auth.AuthMiddleware(http.HandlerFunc(battleHandler.SubmitBattleHandler))).Methods("POST")
	r.Handle("/leaderboard", auth.AuthMiddleware(http.HandlerFunc(leaderboardHandler.GetLeaderboardHandler))).Methods("GET")

	// Run battle processor in a separate Goroutine
	go battleService.ProcessBattlesConcurrently(ctx)

	// Start the server
	http.ListenAndServe(":8080", r)
}
