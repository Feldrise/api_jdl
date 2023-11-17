package models

import "time"

type Group struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of group
	CreatedAt time.Time  `json:"created_at"`            // creation time of group
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of group
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of group

	Code string `json:"code"` // the code to join the group
	Name string `json:"name"` // the name of the group
} // @name Group
