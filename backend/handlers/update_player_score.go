package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"SnakeGame/backend/models"

	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v4"
)

func getAndVerifyUpdateScoreParams(dbConn *pgx.Conn, r *http.Request, player *models.PlayerModel) error {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
			return err
	}

	player.PlayerID = int(claims["player_id"].(float64)) + 1

	query := "SELECT username FROM players WHERE player_id = $1;"
	row := dbConn.QueryRow(context.Background(), query, player.PlayerID)
	err = row.Scan(player.Username)
	if err == pgx.ErrNoRows {
			return ErrInvalidPlayer
	}

	err = json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
			return ErrInvalidJSON
	}

	if player.MaxScore == 0 {
			return ErrMissingMaxScore
	}

	return nil
}
