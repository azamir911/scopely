package services

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"

	"battles/pkg/models"
)

var playerLocks = make(map[string]*sync.Mutex)
var lockMutex = sync.Mutex{}

func ProcessBattlesConcurrently(redisClient *redis.Client) {
	ctx := context.Background()

	for {
		battleJSON, err := redisClient.BLPop(ctx, 0, "battleQueue").Result()
		if err != nil {
			log.Printf("Error popping from battle queue: %v", err)
			continue
		}

		var battle BattleRequest
		if err := json.Unmarshal([]byte(battleJSON[1]), &battle); err != nil {
			log.Printf("Error unmarshaling battle request: %v", err)
			continue
		}

		// Lock both players to prevent conflicting battles
		lockPlayers(battle.AttackerID, battle.DefenderID)

		go func(b BattleRequest) {
			defer unlockPlayers(b.AttackerID, b.DefenderID)
			processBattleRequest(ctx, redisClient, b)
		}(battle)
	}
}

func processBattleRequest(ctx context.Context, redisClient *redis.Client, battle BattleRequest) {
	attackerKey := "player:" + battle.AttackerID
	defenderKey := "player:" + battle.DefenderID

	attackerJSON, err := redisClient.Get(ctx, attackerKey).Result()
	if err != nil {
		log.Printf("Error fetching attacker from Redis: %v", err)
		return
	}

	defenderJSON, err := redisClient.Get(ctx, defenderKey).Result()
	if err != nil {
		log.Printf("Error fetching defender from Redis: %v", err)
		return
	}

	var attacker models.Player
	var defender models.Player

	if err := json.Unmarshal([]byte(attackerJSON), &attacker); err != nil {
		log.Printf("Error unmarshaling attacker: %v", err)
		return
	}

	if err := json.Unmarshal([]byte(defenderJSON), &defender); err != nil {
		log.Printf("Error unmarshaling defender: %v", err)
		return
	}

	// Log battle start
	log.Printf("Processing battle between %s and %s", attacker.Name, defender.Name)

	// Turn-based battle system
	winner, loser := executeBattle(ctx, &attacker, &defender)

	// Update the players
	updatePlayerResources(ctx, redisClient, winner, loser)

	// Log battle outcome
	log.Printf("Battle result: Winner - %s, Loser - %s", winner.Name, loser.Name)

	// Update the leaderboard
	updateLeaderboard(ctx, redisClient, winner, loser)

	// Log leaderboard update
	log.Printf("Leaderboard updated: Winner - %s, Loser - %s", winner.ID, loser.ID)
}

func lockPlayers(attackerID, defenderID string) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	if _, exists := playerLocks[attackerID]; !exists {
		playerLocks[attackerID] = &sync.Mutex{}
	}
	if _, exists := playerLocks[defenderID]; !exists {
		playerLocks[defenderID] = &sync.Mutex{}
	}

	playerLocks[attackerID].Lock()
	playerLocks[defenderID].Lock()
}

func unlockPlayers(attackerID, defenderID string) {
	playerLocks[attackerID].Unlock()
	playerLocks[defenderID].Unlock()
}
