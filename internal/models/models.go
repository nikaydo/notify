package models

import "github.com/google/uuid"

type Event struct {
	Uuid           uuid.UUID
	Title          string
	Description    string
	MaxMembers     int
	CurrectMembers int
	Members        []uuid.UUID
}

type User struct {
	Uuid     uuid.UUID `json:"uuid,omitempty"`
	Login    string    `json:"login" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
