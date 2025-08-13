package services

import (
	"errors"
	"fmt"
	"smart-command-center-backend/config"
	"smart-command-center-backend/firebase"
	"smart-command-center-backend/models"
	"time"
)

// AmbulanceLocationRequest represents the payload for updating an ambulance location.
type AmbulanceLocationRequest struct {
	AmbulanceID uint    `json:"ambulance_id" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
}

// UpdateAmbulanceLocation inserts or updates the latest location of an ambulance.
func UpdateAmbulanceLocation(input AmbulanceLocationRequest) (*models.AmbulanceLocation, error) {
	if input.AmbulanceID == 0 {
		return nil, errors.New("ambulance ID is required")
	}

	var location models.AmbulanceLocation
	err := config.DB.Where("ambulance_id = ?", input.AmbulanceID).First(&location).Error
	if err != nil {
		// Create new record if not found
		location = models.AmbulanceLocation{
			AmbulanceID: input.AmbulanceID,
			Latitude:    input.Latitude,
			Longitude:   input.Longitude,
			UpdatedAt:   time.Now(),
		}
		if err := config.DB.Create(&location).Error; err != nil {
			return nil, err
		}
		// Notify subscribers about new ambulance location
		_ = firebase.SendNotificationToTopic("ambulance_tracking", "Ambulance Tracking", fmt.Sprintf("Ambulance %d location updated", input.AmbulanceID))
		return &location, nil
	}

	// Update existing record
	location.Latitude = input.Latitude
	location.Longitude = input.Longitude
	location.UpdatedAt = time.Now()
	if err := config.DB.Save(&location).Error; err != nil {
		return nil, err
	}

	// Notify subscribers about ambulance movement
	_ = firebase.SendNotificationToTopic("ambulance_tracking", "Ambulance Tracking", fmt.Sprintf("Ambulance %d location updated", input.AmbulanceID))

	return &location, nil
}
