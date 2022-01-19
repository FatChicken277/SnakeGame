package handlers

import (
	"io/ioutil"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func TestResponseOK(t *testing.T) {
	c := require.New(t)
	w := httptest.NewRecorder()

	resp := ResponseOK
	resp.Message = "example"
	resp.Token = "example"

	err := NewResponse(resp, w)
	c.NoError(err)

	wResp := w.Result()
	c.Equal(http.StatusOK, wResp.StatusCode)

	body, err := ioutil.ReadAll(wResp.Body)
	c.NoError(err)
	expected := `{"message":"example","token":"example"}`
	c.Equal(expected, string(body))
}

func TestResponseCreated(t *testing.T) {
	c := require.New(t)
	w := httptest.NewRecorder()

	resp := ResponseCreated
	resp.Data = []int{1, 2, 3, 4}

	err := NewResponse(resp, w)
	c.NoError(err)

	wResp := w.Result()
	c.Equal(http.StatusCreated, wResp.StatusCode)

	body, err := ioutil.ReadAll(wResp.Body)
	c.NoError(err)
	expected := `{"data":[1,2,3,4]}`
	c.Equal(expected, string(body))
}
