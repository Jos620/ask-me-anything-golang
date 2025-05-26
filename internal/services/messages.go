package services

import (
	"github.com/Jos620/ask-me-anything-golang/internal/models"
	"github.com/google/uuid"
)

type MessagesServiceDatabase interface {
	models.MessageRepository
	models.RoomRepository
}

type MessagesService struct {
	db MessagesServiceDatabase
}

func NewMessagesService(db MessagesServiceDatabase) *MessagesService {
	return &MessagesService{
		db: db,
	}
}

func (s *MessagesService) CreateMessage(roomID uuid.UUID, message string) (models.Message, error) {
	// TODO: validate message

	// Validate if the room exists
	_, err := s.db.GetRoomByID(roomID)
	if err != nil {
		return models.Message{}, err
	}

	// Create the message
	messageEntity, err := s.db.CreateMessage(roomID, message)
	if err != nil {
		return models.Message{}, err
	}

	return messageEntity, nil
}

func (s *MessagesService) ReactToMessage(messageID uuid.UUID) (models.Message, error) {
	// Validate if the message exists
	_, err := s.db.GetMessageByID(messageID)
	if err != nil {
		return models.Message{}, err
	}

	// React to the message
	message, err := s.db.ReactToMessage(messageID)
	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}

func (s *MessagesService) RemoveReactionFromMessage(messageID uuid.UUID) (models.Message, error) {
	// Validate if the message exists
	_, err := s.db.GetMessageByID(messageID)
	if err != nil {
		return models.Message{}, err
	}

	// Remove react from message
	message, err := s.db.RemoveReactionFromMessage(messageID)
	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}

func (s *MessagesService) MarkMessageAsAnswered(messageID uuid.UUID) (models.Message, error) {
	// Validate if the message exists
	_, err := s.db.GetMessageByID(messageID)
	if err != nil {
		return models.Message{}, err
	}

	// Mark the message as answered
	message, err := s.db.MarkMessageAsAnswered(messageID)
	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}
