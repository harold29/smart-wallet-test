package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID          string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
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
