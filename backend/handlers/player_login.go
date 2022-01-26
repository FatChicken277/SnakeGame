package handlers

import (
	"SnakeGame/backend/authentication"
	"SnakeGame/backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/jackc/pgx/v4"
)

func getAndVerifyLoginParams(dbConn *pgx.Conn, r *http.Request, player *models.PlayerModel) error {
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		return ErrInvalidJSON
	}

	if player.Username == "" {
		return ErrMissingUsername
	}

	if player.Password == "" {
		return ErrMissingPassword
	}

	return nil
}

// PlayerLogin is in charge of a player login
func PlayerLogin(dbConn *pgx.Conn, tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player models.PlayerModel

		err := getAndVerifyLoginParams(dbConn, r, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseBadRequest, w, err.Error()))
			return
		}

		err = authentication.PlayerAuthVerification(dbConn, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseUnauthorized, w, err.Error()))
			return
		}

		_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{"player_id": strconv.Itoa(player.PlayerID)})
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseInternalServerError, w, err.Error()))
			return
		}

		resp := ResponseOK
		resp.Token = tokenString

		LogError(NewResponse(resp, w))
	}
}
