package services

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

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

// Create a new random source and generator
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func executeBattle(_ context.Context, attacker, defender *models.Player) (*models.Player, *models.Player) {
	for attacker.HitPoints > 0 && defender.HitPoints > 0 {
		// Use defender's luck value to determine if the attack misses
		if rnd.Float64() < defender.LuckValue {
			log.Printf("%s's attack missed!", attacker.Name)
		} else {
			// Calculate damage and apply to defender
			damage := calculateDamage(attacker)
			defender.HitPoints -= damage
			log.Printf("%s attacks %s for %d damage", attacker.Name, defender.Name, damage)
		}

		// Swap roles: attacker becomes defender, and defender becomes attacker
		attacker, defender = defender, attacker
	}

	// Determine winner and loser
	if attacker.HitPoints > 0 {
		return attacker, defender
	}
	return defender, attacker
}

func calculateDamage(player *models.Player) int {
	// Calculate health percentage
	healthPercentage := float64(player.HitPoints) / float64(100)

	// Calculate effective attack value, considering current health percentage
	effectiveAttack := float64(player.AttackValue) * healthPercentage

	// Cap the damage to at least half of the base attack value, as specified
	damage := int(effectiveAttack)
	if damage < player.AttackValue/2 {
		damage = player.AttackValue / 2
	}

	// Ensure minimum damage is at least 1
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
