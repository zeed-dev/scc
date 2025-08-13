package services

import (
	"errors"
	"smart-command-center-backend/config"
	"smart-command-center-backend/firebase"
	"smart-command-center-backend/models"
	"time"
)

type PanicRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	AudioURL  string  `json:"audio_url"` // Optional field for audio recording URL
}

func SendPanic(userID uint, input PanicRequest) (*models.PanicEvent, error) {
	if input.Latitude == 0 || input.Longitude == 0 {
		return nil, errors.New("latitude and longitude are required")
	}

	panicData := models.PanicEvent{
		UserID:    userID,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		AudioURL:  input.AudioURL,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&panicData).Error; err != nil {
		return nil, errors.New("failed to create panic event")
	}

	// Send notification to responders
	_ = firebase.SendNotificationToTopic("panic_alerts", "Panic Alert", "A new panic event has been reported")

	return &panicData, nil
}
