package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Geolocation struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Latitude   string    `gorm:"not null"`
	Longitude  string    `gorm:"not null"`
	LocationID string    `gorm:"not null"`
}
