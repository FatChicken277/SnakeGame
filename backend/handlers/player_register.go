package handlers

import (
	"SnakeGame/backend/models"
	"SnakeGame/backend/storage"
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func getAndVerifyRegisterParams(dbConn *pgx.Conn, r *http.Request, player *models.PlayerModel) error {
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		return ErrInvalidJSON
	}

	if player.Username == "" {
		return ErrMissingUsername
	}

	if len(player.Username) > 30 {
		return ErrInvalidUsernameLength
	}

	query := "SELECT player_id FROM players WHERE username = $1;"
	row := dbConn.QueryRow(context.Background(), query, player.Username)
	err = row.Scan(player.PlayerID)
	if err != pgx.ErrNoRows {
		return ErrExistingUsername
	}

	if player.Password == "" {
		return ErrMissingPassword
	}

	if len(player.Password) < 6 {
		return ErrInvalidPasswordLength
	}

	if player.PasswordConfirm == "" {
		return ErrMissingPasswordConfirm
	}

	if player.Password != player.PasswordConfirm {
		return ErrPassworMatch
	}

	return nil
}

// PlayerRegister is in charge of registering and storing new players
func PlayerRegister(dbConn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player models.PlayerModel

		err := getAndVerifyRegisterParams(dbConn, r, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseBadRequest, w, err.Error()))
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseInternalServerError, w, err.Error()))
			return
		}
		player.PasswordHash = string(passwordHash)

		err = storage.AddPlayer(dbConn, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseInternalServerError, w, err.Error()))
			return
		}

		resp := ResponseCreated
		resp.Message = "the player was successfully created"

		LogError(NewResponse(resp, w))
	}
}
