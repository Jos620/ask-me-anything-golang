package models

import "github.com/google/uuid"

type Message struct {
	ID            uuid.UUID `json:"id"`
	RoomID        uuid.UUID `json:"room_id"`
	Message       string    `json:"message"`
	ReactionCount int64     `json:"reaction_count"`
	Answered      bool      `json:"answered"`
}

type MessageRepository interface {
	GetAllMessages() []Message
	GetMessageByID(id uuid.UUID) (Message, error)
	GetMessagesByRoomID(roomID uuid.UUID) ([]Message, error)
	CreateMessage(roomID uuid.UUID, message string) (Message, error)
	DeleteMessage(id uuid.UUID) error
	UpdateMessage(id uuid.UUID, message string) (Message, error)
	ReactToMessage(id uuid.UUID) (Message, error)
	RemoveReactionFromMessage(id uuid.UUID) (Message, error)
	MarkMessageAsAnswered(id uuid.UUID) (Message, error)
}
