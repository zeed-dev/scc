package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// RouteRequest holds the coordinates for a routing request.
type RouteRequest struct {
	OriginLat float64 `form:"origin_lat" binding:"required"`
	OriginLng float64 `form:"origin_lng" binding:"required"`
	DestLat   float64 `form:"dest_lat" binding:"required"`
	DestLng   float64 `form:"dest_lng" binding:"required"`
}

// RouteResponse is a simplified representation of Google Maps Directions API response.
type RouteResponse struct {
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Polyline string `json:"polyline"`
}

// GetRoute calls Google Maps Directions API and returns routing info.
func GetRoute(input RouteRequest) (*RouteResponse, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		return nil, errors.New("google maps API key not configured")
	}

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%f,%f&destination=%f,%f&key=%s",
		input.OriginLat, input.OriginLng, input.DestLat, input.DestLng, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Routes []struct {
			Legs []struct {
				Distance struct {
					Text string `json:"text"`
				} `json:"distance"`
				Duration struct {
					Text string `json:"text"`
				} `json:"duration"`
			} `json:"legs"`
			OverviewPolyline struct {
				Points string `json:"points"`
			} `json:"overview_polyline"`
		} `json:"routes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Routes) == 0 || len(data.Routes[0].Legs) == 0 {
		return nil, errors.New("no routes found")
	}

	route := data.Routes[0]
	leg := route.Legs[0]
	result := RouteResponse{
		Distance: leg.Distance.Text,
		Duration: leg.Duration.Text,
		Polyline: route.OverviewPolyline.Points,
	}
	return &result, nil
}
