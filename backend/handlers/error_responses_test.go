package handlers

import (
	"io/ioutil"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func TestNewInternalServerResponse(t *testing.T) {
	c := require.New(t)
	w := httptest.NewRecorder()

	err := NewErrorResponse(ErrorResponseInternalServerError, w, "example")
	c.NoError(err)

	resp := w.Result()
	c.Equal(http.StatusInternalServerError, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Internal Server Error","message":"example","status_code":500}`
	c.Equal(expected, string(body))
}

func TestNewBadRequestResponse(t *testing.T) {
	c := require.New(t)
	w := httptest.NewRecorder()

	err := NewErrorResponse(ErrorResponseBadRequest, w, "example")
	c.NoError(err)

	resp := w.Result()
	c.Equal(http.StatusBadRequest, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Bad Request","message":"example","status_code":400}`
	c.Equal(expected, string(body))
}

func TestNewUnauthorizedResponse(t *testing.T) {
	c := require.New(t)
	w := httptest.NewRecorder()

	err := NewErrorResponse(ErrorResponseUnauthorized, w, "example")
	c.NoError(err)

	resp := w.Result()
	c.Equal(http.StatusUnauthorized, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	c.NoError(err)
	expected := `{"status":"Unauthorized","message":"example","status_code":401}`
	c.Equal(expected, string(body))
}
