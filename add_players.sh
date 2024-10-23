#!/bin/bash

# Add Player 1 (Warrior)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player1",
    "name": "Warrior",
    "description": "A fierce warrior",
    "gold": 1000,
    "silver": 500,
    "attack_value": 70,
    "hit_points": 100,
    "luck_value": 0.2
}'

# Add Player 2 (Mage)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player2",
    "name": "Mage",
    "description": "A powerful mage",
    "gold": 800,
    "silver": 600,
    "attack_value": 60,
    "hit_points": 100,
    "luck_value": 0.3
}'

# Add Player 3 (Rogue)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player3",
    "name": "Rogue",
    "description": "A stealthy rogue",
    "gold": 600,
    "silver": 400,
    "attack_value": 50,
    "hit_points": 100,
    "luck_value": 0.4
}'

# Add Player 4 (Paladin)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player4",
    "name": "Paladin",
    "description": "A righteous paladin",
    "gold": 1000,
    "silver": 700,
    "attack_value": 65,
    "hit_points": 120,
    "luck_value": 0.2
}'

# Add Player 5 (Archer)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player5",
    "name": "Archer",
    "description": "An agile archer",
    "gold": 500,
    "silver": 400,
    "attack_value": 55,
    "hit_points": 90,
    "luck_value": 0.3
}'

# Add Player 6 (Berserker)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player6",
    "name": "Berserker",
    "description": "A fierce berserker",
    "gold": 700,
    "silver": 300,
    "attack_value": 80,
    "hit_points": 110,
    "luck_value": 0.15
}'

# Add Player 7 (Knight)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player7",
    "name": "Knight",
    "description": "A brave knight",
    "gold": 900,
    "silver": 500,
    "attack_value": 65,
    "hit_points": 110,
    "luck_value": 0.25
}'

# Add Player 8 (Sorcerer)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player8",
    "name": "Sorcerer",
    "description": "A mysterious sorcerer",
    "gold": 600,
    "silver": 700,
    "attack_value": 70,
    "hit_points": 95,
    "luck_value": 0.35
}'

# Add Player 9 (Assassin)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player9",
    "name": "Assassin",
    "description": "A swift assassin",
    "gold": 750,
    "silver": 350,
    "attack_value": 85,
    "hit_points": 80,
    "luck_value": 0.5
}'

# Add Player 10 (Druid)
curl -X POST http://localhost:8080/player -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "id": "player10",
    "name": "Druid",
    "description": "A nature-loving druid",
    "gold": 500,
    "silver": 500,
    "attack_value": 50,
    "hit_points": 100,
    "luck_value": 0.4
}'
