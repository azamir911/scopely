#!/bin/bash

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player4",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player4"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player4",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player4"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player4",
    "defender_id": "player3"
}' &

curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d '{
    "attacker_id": "player2",
    "defender_id": "player4"
}' &

