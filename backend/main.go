package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Firebase authentication and token refreshing
	config.SetupFirebase()

	// Initialize Gin router
	r := gin.Default()

	// Setup API routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}
