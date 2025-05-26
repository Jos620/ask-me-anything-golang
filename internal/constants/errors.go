package constants

import "errors"

// General
var (
	ErrInvalidUUID = errors.New("invalid UUID")
)

// Rooms
var (
	ErrRoomNotFound = errors.New("room not found")
)

// Messages
var (
	ErrMessageNotFound = errors.New("message not found")
)
