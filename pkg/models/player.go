package models

type Player struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Gold        int64   `json:"gold"`
	Silver      int64   `json:"silver"`
	AttackValue int     `json:"attack_value"`
	HitPoints   int     `json:"hit_points"`
	LuckValue   float64 `json:"luck_value"`
}
