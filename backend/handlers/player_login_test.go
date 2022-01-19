package handlers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"snake_game/backend/storage"
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlayerLogin(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "$2a$10$l8.rKqZmp9fgfX5KqOqy3uFO1UrYPrxZO5SFrzU3ykSYVOObz6A0u")
	c.NoError(err)

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	assert.Regexp(t, `"token":".*"`, string(body))

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerLoginBadJSON(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString("bad json"))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"Invalid JSON string","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerLoginMissingUsername(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(`{}`))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'username' cannot be empty","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerLoginMissingPassword(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'password' cannot be empty","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerLoginInvalidUsername(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusUnauthorized, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Unauthorized","message":"invalid username","status_code":401}`
	c.Equal(expected, string(body))
}

func TestPlayerLoginInvalidPassword(t *testing.T) {
	c := require.New(t)

	tokenAuth := jwtauth.New("HS256", []byte("example"), nil)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "password")
	c.NoError(err)

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"invalid_password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerLogin(dbConn, tokenAuth)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusUnauthorized, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Unauthorized","message":"wrong password","status_code":401}`
	c.Equal(expected, string(body))

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}
