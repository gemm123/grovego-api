package models

import (
	"time"

	"github.com/google/uuid"
)

type InputRecommendationRoute struct {
	StartingPoint string `json:"startingPoint"`
	DistanceTypes int    `json:"distanceTypes"`
	Difficulty    int    `json:"difficulty"`
	Scenery       int    `json:"scenery"`
}

type RecommendationRouteFromML struct {
	RouteCoordinates string  `json:"routeCoordinates"`
	Distance         float64 `json:"distance"`
	Duration         float64 `json:"duration"`
	RouteName        string  `json:"routeName"`
}

type Route struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          string    `gorm:"type:uuid"`
	RouteCoordinate string    `json:"routeCoordinate"`
	RouteName       string
	Distance        float64
	Duration        time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type InputFinishRoute struct {
	RouteCoordinate string  `json:"routeCoordinate"`
	RouteName       string  `json:"routeName"`
	Distance        float64 `json:"distance"`
	Duration        string  `json:"duration"`
}
