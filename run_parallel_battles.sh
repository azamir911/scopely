#!/bin/bash

# Submit multiple battle requests concurrently
for i in {1..10}
do
    attacker_id="player$((RANDOM % 10 + 1))"
    defender_id="player$((RANDOM % 10 + 1))"

    # Ensure attacker and defender are different
    while [ "$attacker_id" == "$defender_id" ]
    do
        defender_id="player$((RANDOM % 10 + 1))"
    done

    curl -X POST http://localhost:8080/battle -H "Content-Type: application/json" -H "Authorization: your_secret_token" -d "{
        \"attacker_id\": \"$attacker_id\",
        \"defender_id\": \"$defender_id\"
    }" &
done

# Wait for all background jobs to finish
wait
