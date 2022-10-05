package models

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	ID              string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TransactionType string
	Description     string
}
