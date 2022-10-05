package models

import "gorm.io/gorm"

type Geolocation struct {
	gorm.Model
	ID         string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Latitude   string
	Longitude  string
	LocationID string
}
