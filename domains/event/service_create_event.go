package event

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Uno-count/event-booking-api/webserver/handler/app"
	"github.com/gin-gonic/gin"
)

type services struct {
	AppConfig *app.App
}

func NewService(appConfig *app.App) *services {
	return &services{AppConfig: appConfig}
}

func (s *services) CreateEventService(ctx context.Context, model Event) error {
	db := s.AppConfig.DB_INSTANCE
	tx, txErr := db.Begin()

	if txErr != nil {
		slog.Error("unable to create transaction", "error", txErr)
		return txErr
	}

	defer tx.Rollback()

	sAffected, sError := dbCreateEvent(ctx, tx, model)

	if sError != nil {
		slog.Error("error on creating an event")
		return sError
	}

	if sAffected == 0 {
		slog.Error("no rows affected")
		return sError
	}

	if err := tx.Commit(); err != nil {
		slog.Error("failed to commit transaction", "error", err)
		return err
	}

	return nil
}

func (s *services) CreateEventHandler(ctx *gin.Context) {
	var model Event
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	if err := s.CreateEventService(ctx.Request.Context(), model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}
