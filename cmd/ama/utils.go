package main

import (
	"github.com/Jos620/ask-me-anything-golang/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getParamRoomID(ctx *gin.Context) (uuid.UUID, error) {
	roomID := ctx.Param("room_id")
	parsedRoomID, err := uuid.Parse(roomID)
	if err != nil {
		return parsedRoomID, constants.ErrInvalidUUID
	}

	return parsedRoomID, nil
}

func getParamMessageID(ctx *gin.Context) (uuid.UUID, error) {
	messageID := ctx.Param("message_id")
	parsedMessageID, err := uuid.Parse(messageID)
	if err != nil {
		return parsedMessageID, constants.ErrInvalidUUID
	}

	return parsedMessageID, nil
}
