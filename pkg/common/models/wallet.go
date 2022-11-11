package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	WalletKey  string
	UserID     string
	Currency   Currency
	CurrencyID string
	Available  bool
	Amount     float64
}
