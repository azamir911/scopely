# Battles Game Backend

This project is a backend service for a battles game, implemented using Go. The backend provides features for registering players, submitting battle requests, processing battles, and maintaining a leaderboard.

## Features

- Register new players with different attributes such as gold, silver, attack value, hit points, and luck value.
- Submit battle requests between players.
- Process battles in a turn-based manner with concurrency support.
- Maintain a leaderboard showing player rankings based on their performance.
- Redis is used as the main database for storing players, battles, and the leaderboard.

## Prerequisites

- Go 1.20
- Redis server
- Bash (for running scripts)

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/azamir911/scopely.git
   cd scopely
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run Redis server (if not already running):
   ```bash
   redis-server
   ```

### Running the Application

To run the application, use the following command:

```bash
go run cmd/main.go
```

### API Endpoints

- **Register Player**: `POST /player`
- **Submit Battle**: `POST /battle`
- **Get Leaderboard**: `GET /leaderboard`

### Authorization

All API endpoints require an authorization header (`Authorization: your_secret_token`).

## Bash Scripts

The following bash scripts are provided to help with adding players and running parallel battles:

### Add Players

The script `add_players.sh` can be used to add 10 players to the system. Run the script using:

```bash
bash add_players.sh
```

### Run Parallel Battles

The script `run_parallel_battles.sh` can be used to submit multiple battle requests concurrently to test the system's parallel processing capabilities. Run the script using:

```bash
bash run_parallel_battles.sh
```

## Cleaning Redis

To clear all data in Redis, you can use the following command:

```bash
redis-cli FLUSHALL
```

## License

This project is licensed under the MIT License.

## Contributions

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## Contact

For any questions, please contact [zamir80@gmail.com].

