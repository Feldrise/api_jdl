package models

import (
	"errors"
	"net/http"
	"time"
)

type GameCard struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of game card
	CreatedAt time.Time  `json:"created_at"`            // creation time of game card
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of game card
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of game card

	GameID uint `json:"game_id"` // the id of the game the card belongs to

	Content string `json:"content"` // the card's content
} // @name GameCard

type GameCardPostPutPayload struct {
	Content *string `json:"content" validate:"required"`
} // @name GameCardPostPutPayload

func (g *GameCardPostPutPayload) Bind(r *http.Request) error {
	if g.Content == nil {
		return errors.New("missing required content property")
	}

	return nil
}
