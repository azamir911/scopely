package services

// import (
//
//	"context"
//	"encoding/json"
//	"log"
//	"math/rand"
//
//	"github.com/go-redis/redis/v8"
//
//	"battles/pkg/models"
//
// )
//
//	type BattleRequest struct {
//		AttackerID string `json:"attacker_id"`
//		DefenderID string `json:"defender_id"`
//	}
//
//	func ProcessBattles(redisClient *redis.Client) {
//		ctx := context.Background()
//
//		for {
//			// Blocking pop from battle queue
//			battleJSON, err := redisClient.BLPop(ctx, 0, "battleQueue").Result()
//			if err != nil {
//				log.Printf("Error popping from battle queue: %v", err)
//				continue
//			}
//
//			// Unmarshal the battle request
//			var battle BattleRequest
//			if err := json.Unmarshal([]byte(battleJSON[1]), &battle); err != nil {
//				log.Printf("Error unmarshaling battle request: %v", err)
//				continue
//			}
//
//			processBattleRequest(ctx, redisClient, battle)
//		}
//	}
