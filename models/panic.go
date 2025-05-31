package models

import "time"

type PanicEvent struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	Latitude  float64 `json:"latitude" gorm:"not null"`
	Longitude float64 `json:"longitude" gorm:"not null"`
	AudioURL  string  `json:"audio_url"` // Optional field for audio recording URL
	CreatedAt time.Time
}
