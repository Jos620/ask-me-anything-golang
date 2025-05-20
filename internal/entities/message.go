package entities

import "github.com/google/uuid"

type Message struct {
	ID            uuid.UUID `json:"id"`
	RoomID        uuid.UUID `json:"room_id"`
	Message       string    `json:"message"`
	ReactionCount int64     `json:"reaction_count"`
	Answered      bool      `json:"answered"`
}
