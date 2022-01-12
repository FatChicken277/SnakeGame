package models

import "encoding/json"

// PlayerModel defines the model for players

type PlayerModel struct {
	PlayerID        int    `json:"player_id,omitempty"`
	Username        string `json:"username,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
	PasswordHash    string `json:"-,omitempty"`
	MaxScore        uint   `json:"max_score,omitempty"`
}

// ToResponse is use to convert a player into a json response
func (player *PlayerModel) ToResponse() ([]byte, error) {
	jsonResponse, err := json.Marshal(player)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
