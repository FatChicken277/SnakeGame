package handlers

import (
	"encoding/json"
	"net/http"
)

// Response is represents the http responses
type Response struct {
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Token      interface{} `json:"token,omitempty"`
	StatusCode int         `json:"-"`
}

var (
	// ResponseOK is used to display an status ok response
	ResponseOK = Response{
		StatusCode: http.StatusOK,
	}
	// ResponseCreated is used to display an bad request error
	ResponseCreated = Response{
		StatusCode: http.StatusCreated,
	}
)

// NewResponse is used to display an specific response to the user
func NewResponse(responseType Response, w http.ResponseWriter) error {
	w.WriteHeader(responseType.StatusCode)
	jsonResponse, err := json.Marshal(responseType)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}
