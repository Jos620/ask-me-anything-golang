package entities

import "github.com/google/uuid"

type Room struct {
	ID    uuid.UUID `json:"id"`
	Theme string    `json:"theme"`
}
