package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Uno-count/event-booking-api/domains/event"
	"github.com/Uno-count/event-booking-api/webserver/handler/app"
	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := app.Init()

	error := appConfig.StartWebServer(context.Background())
	if error != nil {
		fmt.Println("failed to start the app: ", error)
	}

	server := gin.Default()
	eventService := event.NewService(appConfig)
	server.GET("/events")
	server.POST("/events", eventService.CreateEventHandler)

	log.Println("Server starting on :8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
