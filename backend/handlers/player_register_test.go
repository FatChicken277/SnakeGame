package handlers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"snake_game/backend/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerRegister(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password","password_confirm":"password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusCreated, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"message":"the player was successfully created"}`
	c.Equal(expected, string(body))

	query := "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerRegisterBadJSON(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString("bad json"))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"Invalid JSON string","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterMissingUsername(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(`{}`))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'username' cannot be empty","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterUsernameLength(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'username' cannot be greater than 30 characters","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterExistingUsername(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "hash_password")
	c.NoError(err)

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password","password_confirm":"password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"a player with that username already exists","status_code":400}`
	c.Equal(expected, string(body))

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerRegisterMissingPassword(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'password' cannot be empty","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterPasswordLength(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"short"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'password' cannot be less than 6 characters","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterPasswordConfirm(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"'password_confirm' cannot be empty","status_code":400}`
	c.Equal(expected, string(body))
}

func TestPlayerRegisterPasswordMatch(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	w := httptest.NewRecorder()
	jsonExample := `{"username":"example","password":"password","password_confirm":"notMatch"}`
	r := &http.Request{Body: ioutil.NopCloser(bytes.NewBufferString(jsonExample))}

	fun := PlayerRegister(dbConn)
	fun(w, r)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"passwords do not match","status_code":400}`
	c.Equal(expected, string(body))
}
