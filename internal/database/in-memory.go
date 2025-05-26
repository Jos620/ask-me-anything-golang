package database

import (
	"slices"

	"github.com/Jos620/ask-me-anything-golang/internal/constants"
	"github.com/Jos620/ask-me-anything-golang/internal/models"
	"github.com/google/uuid"
)

type InMemoryDatabase struct {
	rooms    []models.Room
	messages []models.Message
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		rooms:    []models.Room{},
		messages: []models.Message{},
	}
}

func (db *InMemoryDatabase) Seed() {
	var roomsThemes = []string{
		"General Discussion",
		"Technology",
		"Health and Wellness",
	}
	for _, theme := range roomsThemes {
		room := models.Room{
			ID:    uuid.New(),
			Theme: theme,
		}
		db.rooms = append(db.rooms, room)

		var messages = []string{
			"Hello, welcome to the room!",
			"Feel free to ask any questions.",
		}
		for _, message := range messages {
			db.messages = append(db.messages, models.Message{
				ID:      uuid.New(),
				RoomID:  room.ID,
				Message: message,
			})
		}
	}
}

func (db *InMemoryDatabase) GetAllRooms() ([]models.Room, error) {
	return db.rooms, nil
}

func (db *InMemoryDatabase) GetRoomByID(id uuid.UUID) (models.Room, error) {
	for _, room := range db.rooms {
		if room.ID == id {
			return room, nil
		}
	}

	return models.Room{}, constants.ErrRoomNotFound
}

func (db *InMemoryDatabase) CreateRoom(theme string) (models.Room, error) {
	room := models.Room{
		ID:    uuid.New(),
		Theme: theme,
	}
	db.rooms = append(db.rooms, room)

	return room, nil
}

func (db *InMemoryDatabase) DeleteRoom(id uuid.UUID) error {
	for i, room := range db.rooms {
		if room.ID == id {
			db.rooms = slices.Delete(db.rooms, i, i+1)
			return nil
		}
	}

	return constants.ErrRoomNotFound
}

func (db *InMemoryDatabase) UpdateRoom(id uuid.UUID, theme string) (models.Room, error) {
	for i, room := range db.rooms {
		if room.ID == id {
			newRoom := models.Room{
				ID:    id,
				Theme: theme,
			}
			db.rooms[i] = newRoom

			return newRoom, nil
		}
	}

	return models.Room{}, constants.ErrRoomNotFound
}

func (db *InMemoryDatabase) GetAllMessages() []models.Message {
	return db.messages
}

func (db *InMemoryDatabase) GetMessageByID(id uuid.UUID) (models.Message, error) {
	for _, message := range db.messages {
		if message.ID == id {
			return message, nil
		}
	}

	return models.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) GetMessagesByRoomID(roomID uuid.UUID) ([]models.Message, error) {
	messages := []models.Message{}
	for _, message := range db.messages {
		if message.RoomID == roomID {
			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (db *InMemoryDatabase) CreateMessage(roomID uuid.UUID, message string) (models.Message, error) {
	messageEntity := models.Message{
		ID:      uuid.New(),
		RoomID:  roomID,
		Message: message,
	}
	db.messages = append(db.messages, messageEntity)

	return messageEntity, nil
}

func (db *InMemoryDatabase) DeleteMessage(id uuid.UUID) error {
	for i, message := range db.messages {
		if message.ID == id {
			db.messages = slices.Delete(db.messages, i, i+1)
			return nil
		}
	}

	return constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) UpdateMessage(id uuid.UUID, message string) (models.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			newMessage := models.Message{
				ID:      id,
				Message: message.Message,
			}
			db.messages[i] = newMessage

			return newMessage, nil
		}
	}

	return models.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) ReactToMessage(id uuid.UUID) (models.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.ReactionCount++
			db.messages[i] = message
			return message, nil
		}
	}

	return models.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) RemoveReactionFromMessage(id uuid.UUID) (models.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.ReactionCount--
			db.messages[i] = message
			return message, nil
		}
	}

	return models.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) MarkMessageAsAnswered(id uuid.UUID) (models.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.Answered = true
			db.messages[i] = message
			return message, nil
		}
	}

	return models.Message{}, constants.ErrMessageNotFound
}
