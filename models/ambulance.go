package models

import "time"

// AmbulanceLocation represents the latest known location of an ambulance
// tracked in the system.
type AmbulanceLocation struct {
	ID          uint    `gorm:"primaryKey"`
	AmbulanceID uint    `json:"ambulance_id" gorm:"not null"`
	Latitude    float64 `json:"latitude" gorm:"not null"`
	Longitude   float64 `json:"longitude" gorm:"not null"`
	UpdatedAt   time.Time
}
