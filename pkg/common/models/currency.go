package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Name   string    `gorm:"not null; unique"`
	Symbol string    `gorm:"not null"`
}
