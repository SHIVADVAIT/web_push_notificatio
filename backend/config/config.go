package config

import (
	"backend/utils"
	"log"
)

// ServiceAccountPath stores the path to the Firebase service account key JSON
const ServiceAccountPath = "C:\\Users\\Shivansh sharma\\Downloads\\webpush-51b10-firebase-adminsdk-fbsvc-8316c67cd8.json"

// SetupFirebase initializes Firebase authentication and token refreshing
func SetupFirebase() {
	// Load initial access token
	if err := utils.LoadAccessToken(ServiceAccountPath); err != nil {
		log.Fatalf("❌ Failed to load access token: %v", err)
	}

	// Start auto-refreshing the access token
	utils.RefreshAccessToken(ServiceAccountPath)
}
