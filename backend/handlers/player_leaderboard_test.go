package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"snake_game/backend/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerLeaderboard(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password, max_score) VALUES ($1, $2, $3), ($4, $2, $5);"
	_, err = dbConn.Exec(context.Background(), query, "user", "password", 30, "user2", 31)
	c.NoError(err)

	w := httptest.NewRecorder()
	r := &http.Request{}

	fun := PlayerLeaderboard(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"data":[{"username":"user2","max_score":31},{"username":"user","max_score":30}]}`
	c.Equal(expected, string(body))

	query = "DELETE FROM players WHERE username = 'user' OR username = 'user2';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerLeaderboardBadConnection(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/fail")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	r := &http.Request{}

	fun := PlayerLeaderboard(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusInternalServerError, resp.StatusCode)
}
