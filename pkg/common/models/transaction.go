package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID                  string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TransactionType     TransactionType
	RequesterUserID     string
	OriginWalletID      string
	DestinationWalletID string
	TransactionAmount   float64
	TransactionState    string
}
