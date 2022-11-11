package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	TransactionTypeID   string    `gorm:"not null"`
	TransactionType     TransactionType
	RequesterUserID     string  `gorm:"not null"`
	OriginWalletID      string  `gorm:"not null"`
	DestinationWalletID string  `gorm:"not null"`
	TransactionAmount   float64 `gorm:"not null"`
	TransactionState    string  `gorm:"not null"`
}
