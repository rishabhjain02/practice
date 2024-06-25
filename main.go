package main

import (
	"github.com/rishabh/practice/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabh/practice/repository"
	"github.com/rishabh/practice/usecase"
)

// We would like you to create a crypto notification app. The features it should include:

// Create a server which will be able to take in the following rest APIs

// Create a notification. Line items may include current price of BTC, market trade volume, intra day high price, market cap
// Send a notification to an email
// List sent notifications (sent, outstanding, failed etc.)
// Delete a notification

func main() {
	notificationRepo := repository.NewNotificationRepo("db")
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepo)
	router := gin.Default()
	http.NewNotificationHandler(router, *notificationUsecase)

	router.Run(":8080")
}
