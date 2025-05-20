package repositories

import (
	"github.com/Jos620/ask-me-anything-golang/internal/entities"
	"github.com/google/uuid"
)

type RoomRepository interface {
	GetAllRooms() ([]entities.Room, error)
	GetRoomByID(id uuid.UUID) (entities.Room, error)
	CreateRoom(theme string) (entities.Room, error)
	DeleteRoom(id uuid.UUID) error
	UpdateRoom(id uuid.UUID, theme string) (entities.Room, error)
}
