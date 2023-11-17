package models

import (
	"errors"
	"net/http"
	"time"
)

type Game struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of game
	CreatedAt time.Time  `json:"created_at"`            // creation time of game
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of game
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of game

	GroupID uint `json:"group_id"` // the id of the group the game belongs to

	Name      string     `json:"name"`                                                             // the game's name
	Type      string     `json:"type"`                                                             // the game's type
	GameCards []GameCard `json:"game_cards"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,"` // the game's cards
	GameModes []GameMode `json:"game_modes"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,"` // the game's modes
}

type GamePostPayload struct {
	Name *string `json:"name" validate:"required"` // name of the game
	Type *string `json:"type" validate:"required"` // the type of the game
} // @name GamePostPayload

func (g *GamePostPayload) Bind(r *http.Request) error {
	if g.Name == nil {
		return errors.New("missing required name property")
	}

	if g.Type == nil {
		return errors.New("missing required type property")
	}

	if *g.Type != "cards" {
		return errors.New("type is not a valide type")
	}

	return nil
}
