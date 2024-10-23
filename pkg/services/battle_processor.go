package services

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/go-redis/redis/v8"

	"battles/pkg/models"
)

type BattleRequest struct {
	AttackerID string `json:"attacker_id"`
	DefenderID string `json:"defender_id"`
}

func ProcessBattles(redisClient *redis.Client) {
	ctx := context.Background()

	for {
		// Blocking pop from battle queue
		battleJSON, err := redisClient.BLPop(ctx, 0, "battleQueue").Result()
		if err != nil {
			log.Printf("Error popping from battle queue: %v", err)
			continue
		}

		// Unmarshal the battle request
		var battle BattleRequest
		if err := json.Unmarshal([]byte(battleJSON[1]), &battle); err != nil {
			log.Printf("Error unmarshaling battle request: %v", err)
			continue
		}

		processBattleRequest(ctx, redisClient, battle)
	}
}

func executeBattle(_ context.Context, attacker, defender *models.Player) (*models.Player, *models.Player) {
	for {
		if rand.Float64() > defender.LuckValue {
			damage := calculateDamage(attacker)
			defender.HitPoints -= damage
			if defender.HitPoints <= 0 {
				return attacker, defender
			}
		}

		// Swap attacker and defender for the next turn
		attacker, defender = defender, attacker
	}
}

func calculateDamage(player *models.Player) int {
	healthPercentage := float64(player.HitPoints) / float64(100)
	effectiveAttack := player.AttackValue * int(healthPercentage)
	damage := effectiveAttack / 2
	if damage < 1 {
		damage = 1
	}
	return damage
}

func updatePlayerResources(ctx context.Context, redisClient *redis.Client, winner, loser *models.Player) {
	resourceStealPercentage := 0.1 + rand.Float64()*(0.2-0.1)

	goldStolen := int64(float64(loser.Gold) * resourceStealPercentage)
	silverStolen := int64(float64(loser.Silver) * resourceStealPercentage)

	// Update resources
	loser.Gold -= goldStolen
	loser.Silver -= silverStolen
	winner.Gold += goldStolen
	winner.Silver += silverStolen

	// Save updated players back to Redis
	winnerKey := "player:" + winner.ID
	loserKey := "player:" + loser.ID

	winnerJSON, _ := json.Marshal(winner)
	loserJSON, _ := json.Marshal(loser)

	redisClient.Set(ctx, winnerKey, winnerJSON, 0)
	redisClient.Set(ctx, loserKey, loserJSON, 0)
}

func updateLeaderboard(ctx context.Context, redisClient *redis.Client, winner, loser *models.Player) {
	// Example implementation: Add/update player scores on leaderboard
	redisClient.ZIncrBy(ctx, "leaderboard", 1, winner.ID)
	redisClient.ZIncrBy(ctx, "leaderboard", -1, loser.ID)
}
