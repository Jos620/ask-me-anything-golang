package database

import (
	"slices"

	"github.com/Jos620/ask-me-anything-golang/internal/constants"
	"github.com/Jos620/ask-me-anything-golang/internal/entities"
	"github.com/google/uuid"
)

type InMemoryDatabase struct {
	rooms    []entities.Room
	messages []entities.Message
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		rooms:    []entities.Room{},
		messages: []entities.Message{},
	}
}

func (db *InMemoryDatabase) Seed() {
	var roomsThemes = []string{
		"General Discussion",
		"Technology",
		"Health and Wellness",
	}
	for _, theme := range roomsThemes {
		room := entities.Room{
			ID:    uuid.New(),
			Theme: theme,
		}
		db.rooms = append(db.rooms, room)

		var messages = []string{
			"Hello, welcome to the room!",
			"Feel free to ask any questions.",
		}
		for _, message := range messages {
			db.messages = append(db.messages, entities.Message{
				ID:      uuid.New(),
				RoomID:  room.ID,
				Message: message,
			})
		}
	}
}

func (db *InMemoryDatabase) GetAllRooms() ([]entities.Room, error) {
	return db.rooms, nil
}

func (db *InMemoryDatabase) GetRoomByID(id uuid.UUID) (entities.Room, error) {
	for _, room := range db.rooms {
		if room.ID == id {
			return room, nil
		}
	}

	return entities.Room{}, constants.ErrRoomNotFound
}

func (db *InMemoryDatabase) CreateRoom(theme string) (entities.Room, error) {
	room := entities.Room{
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

func (db *InMemoryDatabase) UpdateRoom(id uuid.UUID, theme string) (entities.Room, error) {
	for i, room := range db.rooms {
		if room.ID == id {
			newRoom := entities.Room{
				ID:    id,
				Theme: theme,
			}
			db.rooms[i] = newRoom

			return newRoom, nil
		}
	}

	return entities.Room{}, constants.ErrRoomNotFound
}

func (db *InMemoryDatabase) GetAllMessages() []entities.Message {
	return db.messages
}

func (db *InMemoryDatabase) GetMessageByID(id uuid.UUID) (entities.Message, error) {
	for _, message := range db.messages {
		if message.ID == id {
			return message, nil
		}
	}

	return entities.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) GetMessagesByRoomID(roomID uuid.UUID) ([]entities.Message, error) {
	messages := []entities.Message{}
	for _, message := range db.messages {
		if message.RoomID == roomID {
			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (db *InMemoryDatabase) CreateMessage(roomID uuid.UUID, message string) (entities.Message, error) {
	messageEntity := entities.Message{
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

func (db *InMemoryDatabase) UpdateMessage(id uuid.UUID, message string) (entities.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			newMessage := entities.Message{
				ID:      id,
				Message: message.Message,
			}
			db.messages[i] = newMessage

			return newMessage, nil
		}
	}

	return entities.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) ReactToMessage(id uuid.UUID) (entities.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.ReactionCount++
			db.messages[i] = message
			return message, nil
		}
	}

	return entities.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) RemoveReactionFromMessage(id uuid.UUID) (entities.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.ReactionCount--
			db.messages[i] = message
			return message, nil
		}
	}

	return entities.Message{}, constants.ErrMessageNotFound
}

func (db *InMemoryDatabase) MarkMessageAsAnswered(id uuid.UUID) (entities.Message, error) {
	for i, message := range db.messages {
		if message.ID == id {
			message.Answered = true
			db.messages[i] = message
			return message, nil
		}
	}

	return entities.Message{}, constants.ErrMessageNotFound
}
