package pgdb

import "github.com/Jos620/ask-me-anything-golang/internal/models"

func roomToModel(r Room) models.Room {
	return models.Room{
		ID:    r.ID,
		Theme: r.Theme,
	}
}

func roomFromModel(r models.Room) Room {
	return Room{
		ID:    r.ID,
		Theme: r.Theme,
	}
}

func messageToModel(m Message) models.Message {
	return models.Message{
		ID:            m.ID,
		RoomID:        m.RoomID,
		Message:       m.Message,
		ReactionCount: m.ReactionCount,
		Answered:      m.Answered,
	}
}

func messageFromModel(m models.Message) Message {
	return Message{
		ID:            m.ID,
		RoomID:        m.RoomID,
		Message:       m.Message,
		ReactionCount: m.ReactionCount,
		Answered:      m.Answered,
	}
}
