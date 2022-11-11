package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	UserID      string
	User        User
	Address1    string
	Address2    string
	Country     string
	CountryCode string
	State       string
	StateCode   string
	ZipCode     int
	Geolocation Geolocation
}
