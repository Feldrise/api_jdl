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

	Content   string     `json:"content"`                                                                             // the card's content
	GameModes []GameMode `json:"modes" gorm:"many2many:gamecard_modes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // the card's mode
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

type GameCardModeAssociationPayload struct {
	ModeID *uint   `json:"mode_id" validate:"required"`
	Type   *string `json:"type" validate:"required"`
} // @name GameCardModeAssociationPayload

func (g *GameCardModeAssociationPayload) Bind(r *http.Request) error {
	if g.ModeID == nil {
		return errors.New("missing required mode_id property")
	}

	if g.Type == nil {
		return errors.New("missing required type property")
	}

	if *g.Type != "add" && *g.Type != "remove" {
		return errors.New("type must be \"add\" or \"remove\"")
	}

	return nil
}
