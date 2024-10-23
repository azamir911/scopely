package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"

	"battles/pkg/models"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

// PushBattle adds a battle request to the battle queue
func (r *RedisRepository) PushBattle(ctx context.Context, battleRequest models.BattleRequest) error {
	battleJSON, err := json.Marshal(battleRequest)
	if err != nil {
		log.Printf("Failed to marshal battle request: %v", err)
		return err
	}
	return r.client.RPush(ctx, "battleQueue", battleJSON).Err()
}

// GetBattle pops a battle request from the battle queue
func (r *RedisRepository) GetBattle(ctx context.Context) (models.BattleRequest, error) {
	result, err := r.client.BLPop(ctx, 0, "battleQueue").Result()
	if err != nil {
		log.Printf("Error popping from battle queue: %v", err)
		return models.BattleRequest{}, err
	}

	var battle models.BattleRequest
	if err := json.Unmarshal([]byte(result[1]), &battle); err != nil {
		log.Printf("Error unmarshaling battle request: %v", err)
		return models.BattleRequest{}, err
	}
	return battle, nil
}

// PushPlayer adds a player to Redis
func (r *RedisRepository) PushPlayer(ctx context.Context, player *models.Player) error {
	playerJSON, err := json.Marshal(player)
	if err != nil {
		log.Printf("Failed to marshal player: %v", err)
		return err
	}
	return r.client.Set(ctx, "player:"+player.ID, playerJSON, 0).Err()
}

// GetLeaderboard retrieves the leaderboard
func (r *RedisRepository) GetLeaderboard(ctx context.Context) ([]models.LeaderboardEntry, error) {
	rawLeaderboard, err := r.client.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Printf("Failed to retrieve leaderboard: %v", err)
		return nil, err
	}

	var leaderboard []models.LeaderboardEntry
	for rank, entry := range rawLeaderboard {
		leaderboard = append(leaderboard, models.LeaderboardEntry{
			Rank:     rank + 1,
			PlayerID: entry.Member.(string),
			Score:    entry.Score,
		})
	}
	return leaderboard, nil
}

// GetPlayer retrieves a player from Redis
func (r *RedisRepository) GetPlayer(ctx context.Context, playerKey string) (models.Player, error) {
	playerJSON, err := r.client.Get(ctx, playerKey).Result()
	if err == redis.Nil {
		return models.Player{}, errors.New("player not found")
	} else if err != nil {
		log.Printf("Error fetching player from Redis: %v", err)
		return models.Player{}, err
	}

	var player models.Player
	if err := json.Unmarshal([]byte(playerJSON), &player); err != nil {
		log.Printf("Error unmarshaling player: %v", err)
		return models.Player{}, err
	}
	return player, nil
}

// UpdateLeaderboard updates the score of a player in the leaderboard
func (r *RedisRepository) UpdateLeaderboard(ctx context.Context, player *models.Player, scoreDelta float64) error {
	err := r.client.ZIncrBy(ctx, "leaderboard", scoreDelta, player.ID).Err()
	if err != nil {
		log.Printf("Failed to update leaderboard for player %s: %v", player.ID, err)
		return err
	}
	return nil
}
