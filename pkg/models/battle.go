package models

type BattleRequest struct {
	AttackerID string `json:"attacker_id"`
	DefenderID string `json:"defender_id"`
}
