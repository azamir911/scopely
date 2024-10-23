package models

type LeaderboardEntry struct {
	Rank     int     `json:"rank"`      // Player's rank position on the leaderboard
	PlayerID string  `json:"player_id"` // Unique identifier of the player
	Score    float64 `json:"score"`     // Player's score on the leaderboard
}
