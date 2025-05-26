package entities

import "github.com/google/uuid"

type Room struct {
	ID    uuid.UUID `json:"id"`
	Theme string    `json:"theme"`
}

type RoomRepository interface {
	GetAllRooms() ([]Room, error)
	GetRoomByID(id uuid.UUID) (Room, error)
	CreateRoom(theme string) (Room, error)
	DeleteRoom(id uuid.UUID) error
	UpdateRoom(id uuid.UUID, theme string) (Room, error)
}
