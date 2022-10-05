package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	ID     string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name   string
	Symbol string
}
