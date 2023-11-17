package models

import "time"

type ModeCard struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of mode card
	CreatedAt time.Time  `json:"created_at"`            // creation time of mode card
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of mode card
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of mode card

	GameModeID uint `json:"game_mode_id"` // the ID of the mode the card belongs to
	CardID     uint `json:"card_id"`      // th ID of the corresponding game card
}
