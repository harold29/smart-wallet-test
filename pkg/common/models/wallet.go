package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID        string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	WalletKey string
	UserID    string
	Currency  Currency
	Available bool
	Amount    float64
}
