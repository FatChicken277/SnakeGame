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
