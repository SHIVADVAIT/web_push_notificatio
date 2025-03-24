package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"golang.org/x/oauth2/google"
)

var (
	accessToken string
	mu          sync.Mutex
)

// LoadAccessToken generates an OAuth 2.0 access token
func LoadAccessToken(serviceAccountPath string) error {
	data, err := ioutil.ReadFile(serviceAccountPath)
	if err != nil {
		return fmt.Errorf("failed to read service account file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return fmt.Errorf("failed to parse credentials: %v", err)
	}

	token, err := config.TokenSource(context.Background()).Token()
	if err != nil {
		return fmt.Errorf("failed to generate access token: %v", err)
	}

	mu.Lock()
	accessToken = token.AccessToken
	mu.Unlock()

	log.Println("New Access Token Generated")
	return nil
}

// RefreshAccessToken refreshes the token every 59 minutes
func RefreshAccessToken(serviceAccountPath string) {
	go func() {
		for {
			if err := LoadAccessToken(serviceAccountPath); err != nil {
				log.Println("Error refreshing access token:", err)
			}
			time.Sleep(59 * time.Minute)
		}
	}()
}

// GetAccessToken returns the current access token
func GetAccessToken() string {
	mu.Lock()
	defer mu.Unlock()
	return accessToken
}
