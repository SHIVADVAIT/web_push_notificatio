package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/send-notification", handlers.SendNotificationHandler)
}
