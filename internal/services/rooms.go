package services

import (
	"github.com/Jos620/ask-me-anything-golang/internal/models"
	"github.com/google/uuid"
)

type RoomsServiceDatabase interface {
	models.RoomRepository
	models.MessageRepository
}

type RoomsService struct {
	db RoomsServiceDatabase
}

func NewRoomsService(db RoomsServiceDatabase) *RoomsService {
	return &RoomsService{
		db: db,
	}
}

func (s *RoomsService) CreateRoom(theme string) (models.Room, error) {
	// TODO: validate theme
	return s.db.CreateRoom(theme)
}

func (s *RoomsService) GetAllRooms() ([]models.Room, error) {
	rooms, err := s.db.GetAllRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *RoomsService) GetRoomMessages(roomID uuid.UUID) ([]models.Message, error) {
	// Validate if the room exists
	_, err := s.db.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	// Get all the messages inside a room
	messages, err := s.db.GetMessagesByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
