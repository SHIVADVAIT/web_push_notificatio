package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"backend/utils"

	"github.com/gin-gonic/gin"
)

// FCMRequest represents the payload for FCM API
type FCMRequest struct {
	Message struct {
		Token        string `json:"token"`
		Notification struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"notification"`
	} `json:"message"`
}

// SendNotificationHandler sends a push notification via FCM
func SendNotificationHandler(c *gin.Context) {
	var req FCMRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Get the latest access token
	accessToken := utils.GetAccessToken()
	if accessToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Access token not available"})
		return
	}

	// FCM API URL
	url := "https://fcm.googleapis.com/v1/projects/webpush-51b10/messages:send"

	// Convert request struct to JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding request"})
		return
	}

	// Create HTTP request
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Set Headers
	httpReq.Header.Set("Authorization", "Bearer "+accessToken)
	httpReq.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request failed"})
		return
	}
	defer resp.Body.Close()

	// Read response
	body, _ := ioutil.ReadAll(resp.Body)

	// Return FCM response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Notification sent successfully",
		"response": string(body),
	})
}
