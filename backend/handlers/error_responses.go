package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse is represents the http error responses
type ErrorResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

var (
	// ErrorResponseInternalServerError is used to display an internal error
	ErrorResponseInternalServerError = ErrorResponse{
		Status:     "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}
	// ErrorResponseBadRequest is used to display an bad request error
	ErrorResponseBadRequest = ErrorResponse{
		Status:     "Bad Request",
		StatusCode: http.StatusBadRequest,
	}
	// ErrorResponseUnauthorized is used to display an unauthorized error
	ErrorResponseUnauthorized = ErrorResponse{
		Status:     "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
)

// LogError is used to log errors
func LogError(err error) {
	if err != nil {
		log.Fatal("Unexpected error: " + err.Error())
	}
}

// NewErrorResponse is used to display an specific error response to the user
func NewErrorResponse(errType ErrorResponse, w http.ResponseWriter, message string) error {
	errType.Message = message

	w.WriteHeader(errType.StatusCode)
	jsonResponse, err := json.Marshal(errType)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}
