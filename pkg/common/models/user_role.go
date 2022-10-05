package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	ID          string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Permission  string
	Description string
}
