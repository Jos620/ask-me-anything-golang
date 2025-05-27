package pgdb

import (
	"context"
	"fmt"
	"os"

	"github.com/Jos620/ask-me-anything-golang/internal/constants"
	"github.com/Jos620/ask-me-anything-golang/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var connection *pgx.Conn

type PostgresDatabase struct {
	queries *Queries
	conn    *pgx.Conn
}

func NewPostgresDatabase() *PostgresDatabase {
	if connection == nil {
		ctx := context.Background()
		conn, err := connectToDatabase(ctx)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to PostgreSQL: %v", err))
		}
		connection = conn
	}

	return &PostgresDatabase{
		queries: New(connection),
		conn:    connection,
	}
}

func connectToDatabase(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("AMA_DATABASE_USER"),
		os.Getenv("AMA_DATABASE_PASSWORD"),
		os.Getenv("AMA_DATABASE_HOST"),
		os.Getenv("AMA_DATABASE_PORT"),
		os.Getenv("AMA_DATABASE_NAME"),
	))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		conn.Close(ctx)
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return conn, nil
}

func (db *PostgresDatabase) Close() error {
	if db.conn != nil {
		return db.conn.Close(context.Background())
	}
	return nil
}

func (db *PostgresDatabase) GetAllRooms() ([]models.Room, error) {
	ctx := context.Background()
	rooms, err := db.queries.GetRooms(ctx)
	if err != nil {
		return nil, err
	}

	modelRooms := make([]models.Room, len(rooms))
	for i, room := range rooms {
		modelRooms[i] = roomToModel(room)
	}

	return modelRooms, nil
}

func (db *PostgresDatabase) GetRoomByID(id uuid.UUID) (models.Room, error) {
	ctx := context.Background()
	room, err := db.queries.GetRoomByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Room{}, constants.ErrRoomNotFound
		}
		return models.Room{}, err
	}

	return roomToModel(room), nil
}

func (db *PostgresDatabase) CreateRoom(theme string) (models.Room, error) {
	ctx := context.Background()
	room, err := db.queries.CreateRoom(ctx, theme)
	if err != nil {
		return models.Room{}, err
	}

	return roomToModel(room), nil
}

func (db *PostgresDatabase) DeleteRoom(id uuid.UUID) error {
	ctx := context.Background()
	err := db.queries.DeleteRoom(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return constants.ErrRoomNotFound
		}
		return err
	}

	return nil
}

func (db *PostgresDatabase) UpdateRoom(id uuid.UUID, theme string) (models.Room, error) {
	ctx := context.Background()
	room, err := db.queries.UpdateRoom(ctx, UpdateRoomParams{
		ID:    id,
		Theme: theme,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Room{}, constants.ErrRoomNotFound
		}
		return models.Room{}, err
	}

	return roomToModel(room), nil
}

func (db *PostgresDatabase) GetAllMessages() []models.Message {
	ctx := context.Background()
	messages, err := db.queries.GetAllMessages(ctx)
	if err != nil {
		return nil
	}

	modelMessages := make([]models.Message, len(messages))
	for i, message := range messages {
		modelMessages[i] = messageToModel(message)
	}

	return modelMessages
}

func (db *PostgresDatabase) GetMessageByID(id uuid.UUID) (models.Message, error) {
	ctx := context.Background()
	message, err := db.queries.GetMessageByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Message{}, constants.ErrMessageNotFound
		}
		return models.Message{}, err
	}

	return messageToModel(message), nil
}

func (db *PostgresDatabase) GetMessagesByRoomID(roomID uuid.UUID) ([]models.Message, error) {
	ctx := context.Background()
	messages, err := db.queries.GetMessagesByRoomID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	modelMessages := make([]models.Message, len(messages))
	for i, message := range messages {
		modelMessages[i] = messageToModel(message)
	}

	return modelMessages, nil
}

func (db *PostgresDatabase) CreateMessage(roomID uuid.UUID, message string) (models.Message, error) {
	ctx := context.Background()
	messageEntity, err := db.queries.CreateMessage(ctx, CreateMessageParams{
		RoomID:  roomID,
		Message: message,
	})
	if err != nil {
		return models.Message{}, err
	}

	return messageToModel(messageEntity), nil
}

func (db *PostgresDatabase) DeleteMessage(id uuid.UUID) error {
	ctx := context.Background()
	err := db.queries.DeleteMessage(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return constants.ErrMessageNotFound
		}
		return err
	}

	return nil
}

func (db *PostgresDatabase) UpdateMessage(id uuid.UUID, message string) (models.Message, error) {
	ctx := context.Background()
	existingMessage, err := db.queries.GetMessageByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Message{}, constants.ErrMessageNotFound
		}
		return models.Message{}, err
	}

	updatedMessage, err := db.queries.UpdateMessage(ctx, UpdateMessageParams{
		ID:            id,
		Message:       message,
		ReactionCount: existingMessage.ReactionCount,
		Answered:      existingMessage.Answered,
	})
	if err != nil {
		return models.Message{}, err
	}

	return messageToModel(updatedMessage), nil
}

func (db *PostgresDatabase) ReactToMessage(id uuid.UUID) (models.Message, error) {
	ctx := context.Background()
	message, err := db.queries.ReactToMessage(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Message{}, constants.ErrMessageNotFound
		}
		return models.Message{}, err
	}

	return messageToModel(message), nil
}

func (db *PostgresDatabase) RemoveReactionFromMessage(id uuid.UUID) (models.Message, error) {
	ctx := context.Background()
	message, err := db.queries.RemoveReactionFromMessage(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Message{}, constants.ErrMessageNotFound
		}
		return models.Message{}, err
	}

	return messageToModel(message), nil
}

func (db *PostgresDatabase) MarkMessageAsAnswered(id uuid.UUID) (models.Message, error) {
	ctx := context.Background()
	message, err := db.queries.MarkMessageAsAnswered(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Message{}, constants.ErrMessageNotFound
		}
		return models.Message{}, err
	}

	return messageToModel(message), nil
}
