package main

import (
	"errors"
	"net/http"

	"github.com/Jos620/ask-me-anything-golang/internal/constants"
	"github.com/Jos620/ask-me-anything-golang/internal/database/imdb"
	"github.com/Jos620/ask-me-anything-golang/internal/services"
	"github.com/gin-gonic/gin"
)

var (
	messagesService services.MessagesService
	roomsService    services.RoomsService
)

func init() {
	inMemoryDatabase := imdb.NewInMemoryDatabase()
	inMemoryDatabase.Seed()

	messagesService = *services.NewMessagesService(inMemoryDatabase)
	roomsService = *services.NewRoomsService(inMemoryDatabase)
}

func main() {
	router := gin.Default()
	apiRouter := router.Group("/api")

	roomsRouter := apiRouter.Group("/rooms")
	roomsRouter.GET("/", handleGetAllRooms)
	roomsRouter.POST("/", handleCreateRoom)

	roomMessagesRouter := roomsRouter.Group(":room_id/messages")
	roomMessagesRouter.GET("/", handleGetRoomMessages)
	roomMessagesRouter.POST("/", handleCreateMessage)
	roomMessagesRouter.PATCH("/answer", handleMarkMessageAsAnswered)

	messageReactionRouter := roomMessagesRouter.Group("/:message_id/react")
	messageReactionRouter.PATCH("/", handleReactToMessage)
	messageReactionRouter.DELETE("/", handleRemoveReactionFromMessage)

	router.Run()
}

func handleGetAllRooms(ctx *gin.Context) {
	rooms, err := roomsService.GetAllRooms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(rooms))
}

func handleCreateRoom(ctx *gin.Context) {
	var requestBody createRoomDTO

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	room, err := roomsService.CreateRoom(requestBody.Theme)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(room))
}

func handleGetRoomMessages(ctx *gin.Context) {
	roomID, err := getParamRoomID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	messages, err := roomsService.GetRoomMessages(roomID)
	if err != nil {
		if errors.Is(err, constants.ErrRoomNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(messages))
}

func handleCreateMessage(ctx *gin.Context) {
	roomID, err := getParamRoomID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var requestBody createMessageDTO
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	message, err := messagesService.CreateMessage(roomID, requestBody.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(message))
}

func handleReactToMessage(ctx *gin.Context) {
	messageID, err := getParamMessageID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	message, err := messagesService.ReactToMessage(messageID)
	if err != nil {
		if errors.Is(err, constants.ErrMessageNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(message))
}

func handleRemoveReactionFromMessage(ctx *gin.Context) {
	messageID, err := getParamMessageID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	message, err := messagesService.RemoveReactionFromMessage(messageID)
	if err != nil {
		if errors.Is(err, constants.ErrMessageNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(message))
}

func handleMarkMessageAsAnswered(ctx *gin.Context) {
	messageID, err := getParamMessageID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	message, err := messagesService.MarkMessageAsAnswered(messageID)
	if err != nil {
		if errors.Is(err, constants.ErrMessageNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, defaultResponse(message))
}
