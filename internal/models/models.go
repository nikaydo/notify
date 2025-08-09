package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Uuid         uuid.UUID   `json:"uuid"`
	Title        string      `json:"title" binding:"required"`
	Description  string      `json:"desctiption" `
	MaxMembers   int         `json:"maxMembers" binding:"required"`
	CountMembers int         `json:"countMembers"`
	Members      []uuid.UUID `json:"members"`
	Created      time.Time   `json:"created"`
	Expaired     time.Time   `json:"expaired" binding:"required"`
	Creator      uuid.UUID   `json:"creator" binding:"required"`
}

type ReturnValues struct {
	UuidEvent uuid.UUID `json:"uuidEvent"`
	UuidUser  uuid.UUID `json:"uuidUser"`
}

type User struct {
	Uuid     uuid.UUID `json:"uuid,omitempty"`
	Login    string    `json:"login" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
