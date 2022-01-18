package storage

import (
	"SnakeGame/backend/models"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDBConection(t *testing.T) {
	c := require.New(t)

	dbSource := "postgresql://admin@localhost:26257/snake_test?sslmode=disable"
	dbConn, err := DBConection(dbSource)
	c.NoError(err)
	dbConn.Close(context.Background())
}

func TestAddPlayer(t *testing.T) {
	c := require.New(t)

	dbSource := "postgresql://admin@localhost:26257/snake_test?sslmode=disable"
	dbConn, err := DBConection(dbSource)
	c.NoError(err)
	defer dbConn.Close(context.Background())

	newPlayer := &models.PlayerModel{
		Username:     "example",
		PasswordHash: "hash_password",
	}

	err = AddPlayer(dbConn, newPlayer)
	c.NoError(err)

	query := "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestUpdatePlayerScore(t *testing.T) {
	c := require.New(t)

	dbSource := "postgresql://admin@localhost:26257/snake_test?sslmode=disable"
	dbConn, err := DBConection(dbSource)
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (player_id, username, password) VALUES ($1, $2, $3);"
	_, err = dbConn.Exec(context.Background(), query, "123", "example", "hash_password")
	c.NoError(err)

	examplePlayer := &models.PlayerModel{
		PlayerID:     123,
		Username:     "example",
		PasswordHash: "hash_password",
		MaxScore:     23,
	}

	err = UpdatePlayerScore(dbConn, examplePlayer)
	c.NoError(err)

	var score uint
	query = "SELECT max_score FROM players WHERE player_id = 123;"
	row := dbConn.QueryRow(context.Background(), query)
	err = row.Scan(&score)
	c.NoError(err)

	c.Equal(score, examplePlayer.MaxScore)

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestGetPlayersLeaderboard(t *testing.T) {
	c := require.New(t)

	dbSource := "postgresql://admin@localhost:26257/snake_test?sslmode=disable"
	dbConn, err := DBConection(dbSource)
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password, max_score) VALUES ($1, $2, $3), ($4, $2, $5);"
	_, err = dbConn.Exec(context.Background(), query, "user", "password", 30, "user2", 31)
	c.NoError(err)

	rows, err := GetPlayersLeaderboard(dbConn)
	c.NoError(err)

	var playerList []models.PlayerModel

	for rows.Next() {
		player := models.PlayerModel{}

		err = rows.Scan(&player.Username, &player.MaxScore)
		c.NoError(err)

		playerList = append(playerList, player)
	}

	c.Equal(uint(31), playerList[0].MaxScore)
	c.Equal(uint(30), playerList[1].MaxScore)

	query = "DELETE FROM players WHERE username = 'user' OR username = 'user2';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}
