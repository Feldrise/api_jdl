package models

import (
	"errors"
	"net/http"
	"time"
)

type GameMode struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of game mode
	CreatedAt time.Time  `json:"created_at"`            // creation time of game mode
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of game mode
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of game mode

	GameID uint `json:"game_id"` // ID of the game the mode belongs to

	Name      string     `json:"name"`                                                             // the mode's name
	ModeCards []ModeCard `json:"mode_cards"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,"` // cards of the mode
}

type GameModePostPutPayload struct {
	Name *string `json:"name" validate:"required"`
}

func (g *GameModePostPutPayload) Bind(r *http.Request) error {
	if g.Name == nil {
		return errors.New("missing required content property")
	}

	return nil
}
