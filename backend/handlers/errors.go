package handlers

import "errors"

var (
	// ErrMissingUsername indicates that the username is missing
	ErrMissingUsername = errors.New("'username' cannot be empty")
	// ErrInvalidUsernameLength indicates that the username is greater than 30 characters
	ErrInvalidUsernameLength = errors.New("'username' cannot be greater than 30 characters")
	// ErrExistingUsername indicates that the username already exists
	ErrExistingUsername = errors.New("a player with that username already exists")
	// ErrMissingPassword indicates that the password is missing
	ErrMissingPassword = errors.New("'password' cannot be empty")
	// ErrInvalidPasswordLength indicates that the password is less than 6 characters
	ErrInvalidPasswordLength = errors.New("'password' cannot be less than 6 characters")
	// ErrMissingPasswordConfirm indicates that the password_confirm is missing
	ErrMissingPasswordConfirm = errors.New("'password_confirm' cannot be empty")
	// ErrPassworMatch indicates that the password and passwordConfirm do not match
	ErrPassworMatch = errors.New("passwords do not match")
	// ErrMissingMaxScore indicates that the max score is missing
	ErrMissingMaxScore = errors.New("'max_score' cannot be empty")

	// ErrInvalidPlayer indicates that the player does not exist
	ErrInvalidPlayer = errors.New("the player associated with the token does not exist")

	// ErrInvalidJSON indicates that an invalid json was entered
	ErrInvalidJSON = errors.New("Invalid JSON string")
)
