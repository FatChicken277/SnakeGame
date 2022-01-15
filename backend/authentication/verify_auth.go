package authentication

import (
	"context"
	"errors"
	"SnakeGame/backend/models"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrInvalidPlayer indicates that the player does not exist
	ErrInvalidPlayer = errors.New("invalid username")
	// ErrInvalidPassword indicates that the password is incorrect
	ErrInvalidPassword = errors.New("wrong password")
)

// PlayerAuthVerification is in charge of verifying if the user is authenticated
func PlayerAuthVerification(conn *pgx.Conn, p *models.PlayerModel) error {
	row := conn.QueryRow(context.Background(), "SELECT player_id, password FROM players WHERE username = $1", p.Username)
	err := row.Scan(&p.PlayerID, &p.PasswordHash)
	if err == pgx.ErrNoRows {
		return ErrInvalidPlayer
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.PasswordHash), []byte(p.Password))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
