package services

import (
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
)

// DashboardStats represents aggregated information for the dashboard.
type DashboardStats struct {
	TotalUsers       int64 `json:"total_users"`
	TotalPanicEvents int64 `json:"total_panic_events"`
	TotalAmbulances  int64 `json:"total_ambulances"`
}

// GetDashboardStats collects simple analytics for the dashboard.
func GetDashboardStats() (*DashboardStats, error) {
	var users, panicEvents, ambulances int64

	if err := config.DB.Model(&models.User{}).Count(&users).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Model(&models.PanicEvent{}).Count(&panicEvents).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Model(&models.AmbulanceLocation{}).Count(&ambulances).Error; err != nil {
		return nil, err
	}

	stats := DashboardStats{
		TotalUsers:       users,
		TotalPanicEvents: panicEvents,
		TotalAmbulances:  ambulances,
	}
	return &stats, nil
}
