package services

import (
	"context"
	"log"
	"math/rand"
	"sync"

	"battles/pkg/models"
	"battles/pkg/repository"
)

type BattleService struct {
	repository *repository.RedisRepository
}

func NewBattleService(repo *repository.RedisRepository) *BattleService {
	return &BattleService{
		repository: repo,
	}
}

func (s *BattleService) ProcessBattlesConcurrently(ctx context.Context) {
	for {
		battle, err := s.repository.GetBattle(ctx)
		if err != nil {
			log.Printf("Error retrieving battle from queue: %v", err)
			continue
		}

		go func(b models.BattleRequest) {
			// Lock players to prevent conflicting battles
			lockPlayers(battle.AttackerID, battle.DefenderID)
			defer unlockPlayers(b.AttackerID, b.DefenderID)
			s.processBattleRequest(ctx, b)
		}(battle)
	}
}

// Locks to ensure battles are non-conflicting (e.g., a player isn't in multiple battles at once)
var playerLocks = make(map[string]*sync.Mutex)
var lockMutex = sync.Mutex{}

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

func (s *BattleService) processBattleRequest(ctx context.Context, battle models.BattleRequest) {
	attackerKey := "player:" + battle.AttackerID
	defenderKey := "player:" + battle.DefenderID

	attacker, err := s.repository.GetPlayer(ctx, attackerKey)
	if err != nil {
		log.Printf("Error fetching attacker from repository: %v", err)
		return
	}

	defender, err := s.repository.GetPlayer(ctx, defenderKey)
	if err != nil {
		log.Printf("Error fetching defender from repository: %v", err)
		return
	}

	log.Printf("Processing battle between %s and %s", attacker.Name, defender.Name)

	// Battle logic
	winner, loser := executeBattle(ctx, &attacker, &defender)

	// Update the players' resources
	s.updatePlayerResources(ctx, winner, loser)

	// Update the leaderboard
	if err := s.updateLeaderboard(ctx, winner, loser); err != nil {
		log.Printf("Error updating leaderboard: %v", err)
	}

	log.Printf("Battle result: Winner - %s, Loser - %s", winner.Name, loser.Name)
}

func (s *BattleService) updateLeaderboard(ctx context.Context, winner, loser *models.Player) error {
	if err := s.repository.UpdateLeaderboard(ctx, winner, 1); err != nil {
		return err
	}
	if err := s.repository.UpdateLeaderboard(ctx, loser, -1); err != nil {
		return err
	}

	// Log leaderboard update
	log.Printf("Leaderboard updated: Winner - %s, Loser - %s", winner.ID, loser.ID)
	return nil
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

func (s *BattleService) updatePlayerResources(ctx context.Context, winner, loser *models.Player) {
	resourceStealPercentage := 0.1 + rand.Float64()*(0.2-0.1)

	goldStolen := int64(float64(loser.Gold) * resourceStealPercentage)
	silverStolen := int64(float64(loser.Silver) * resourceStealPercentage)

	// Update resources
	loser.Gold -= goldStolen
	loser.Silver -= silverStolen
	winner.Gold += goldStolen
	winner.Silver += silverStolen

	// Save updated players back to Redis
	if err := s.repository.PushPlayer(ctx, winner); err != nil {
		log.Printf("Error updating winner: %v", err)
	}
	if err := s.repository.PushPlayer(ctx, loser); err != nil {
		log.Printf("Error updating loser: %v", err)
	}
}
