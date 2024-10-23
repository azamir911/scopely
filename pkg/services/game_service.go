package services

import (
	"context"
	"log"

	"battles/pkg/models"
	"battles/pkg/repository"
)

type GameService struct {
	repository *repository.RedisRepository
}

func NewGameService(repo *repository.RedisRepository) *GameService {
	return &GameService{
		repository: repo,
	}
}

// CreatePlayer creates a new player
func (s *GameService) CreatePlayer(ctx context.Context, player models.Player) error {
	err := s.repository.PushPlayer(ctx, &player)
	if err != nil {
		log.Printf("Failed to create player: %v", err)
		return err
	}
	log.Printf("Player was added: %s", player.ID)
	return nil
}

// SubmitBattle submits a new battle request
func (s *GameService) SubmitBattle(ctx context.Context, battleRequest models.BattleRequest) error {
	err := s.repository.PushBattle(ctx, battleRequest)
	if err != nil {
		log.Printf("Failed to submit battle request: %v", err)
		return err
	}

	log.Printf("Battle request submitted: Attacker - %s, Defender - %s", battleRequest.AttackerID, battleRequest.DefenderID)
	return nil
}

// GetLeaderboard fetches the leaderboard
func (s *GameService) GetLeaderboard(ctx context.Context) ([]models.LeaderboardEntry, error) {
	leaderboard, err := s.repository.GetLeaderboard(ctx)
	if err != nil {
		log.Printf("Failed to retrieve leaderboard: %v", err)
		return nil, err
	}
	return leaderboard, nil
}
