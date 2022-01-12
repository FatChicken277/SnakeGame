package storage

import (
	"context"
	"SnakeGame/backend/models"

	"github.com/jackc/pgx/v4"
)

// DBConection is used to establish connection with the database
func DBConection(source string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(source)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// AddPlayer is used to add a new player into database
func AddPlayer(conn *pgx.Conn, p *models.PlayerModel) error {
	query := "INSERT INTO players (username, password) VALUES ($1, $2);"

	_, err := conn.Exec(context.Background(), query, p.Username, p.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePlayerScore is used to update the player max score
func UpdatePlayerScore(conn *pgx.Conn, p *models.PlayerModel) error {
	query := "UPDATE players SET max_score = $1 WHERE player_id = $2 AND $1 > max_score;"

	_, err := conn.Exec(context.Background(), query, p.MaxScore, p.PlayerID)
	if err != nil {
		return err
	}

	return nil
}

// GetPlayersLeaderboard is used to return the top 10 best scores
func GetPlayersLeaderboard(conn *pgx.Conn) (pgx.Rows, error) {
	query := "SELECT username, max_score FROM players WHERE max_score > 0 ORDER BY max_score DESC LIMIT 10;"

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
