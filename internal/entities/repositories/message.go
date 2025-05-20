package repositories

import (
	"github.com/Jos620/ask-me-anything-golang/internal/entities"
	"github.com/google/uuid"
)

type MessageRepository interface {
	GetAllMessages() []entities.Message
	GetMessageByID(id uuid.UUID) (entities.Message, error)
	GetMessagesByRoomID(roomID uuid.UUID) ([]entities.Message, error)
	CreateMessage(roomID uuid.UUID, message string) (entities.Message, error)
	DeleteMessage(id uuid.UUID) error
	UpdateMessage(id uuid.UUID, message string) (entities.Message, error)
	ReactToMessage(id uuid.UUID) (entities.Message, error)
	RemoveReactionFromMessage(id uuid.UUID) (entities.Message, error)
	MarkMessageAsAnswered(id uuid.UUID) (entities.Message, error)
}
