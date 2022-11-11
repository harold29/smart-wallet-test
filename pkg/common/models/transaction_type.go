package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	TransactionType string    `gorm:"not null"`
	Description     string    `gorm:"not null"`
}
