package event

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *services) GetEventService(ctx context.Context) ([]EventResponse, error) {
	db := s.AppConfig.DB_INSTANCE
	events, err := dbGetEvent(ctx, db)

	if err != nil {
		slog.Error("error on getting events", "error", err)
		return nil, err
	}

	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, EventResponse{
			ID:          event.ID,
			Name:        event.Name,
			Description: event.Description,
			Location:    event.Location,
			UserID:      event.UserID,
		})
	}

	return eventResponses, nil
}

func (s *services) GetEventServiceHandler(ctx *gin.Context) {
	events, err := s.GetEventService(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get events"})
		return
	}

	ctx.JSON(http.StatusOK, events)

}
