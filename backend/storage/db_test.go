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
