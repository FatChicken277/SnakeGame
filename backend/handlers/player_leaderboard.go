package handlers

import (
	"SnakeGame/backend/models"
	"SnakeGame/backend/storage"
	"net/http"

	"github.com/jackc/pgx/v4"
)

// PlayerLeaderboard is used to display the players leaderboard
func PlayerLeaderboard(dbConn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var playerList []models.PlayerModel

		rows, err := storage.GetPlayersLeaderboard(dbConn)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseInternalServerError, w, err.Error()))
			return
		}

		for rows.Next() {
			player := models.PlayerModel{}

			err = rows.Scan(&player.Username, &player.MaxScore)
			if err != nil {
				LogError(NewErrorResponse(ErrorResponseInternalServerError, w, err.Error()))
				return
			}

			playerList = append(playerList, player)
		}

		resp := ResponseOK
		resp.Data = playerList

		LogError(NewResponse(resp, w))
	}
}
