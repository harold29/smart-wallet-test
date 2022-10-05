package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber1 string
	PhoneNumber2 string
	UserTypeID   string
	UserType     UserType
	UserRoleID   string
	UserRole     UserRole
	Wallets      []Wallet
	Transactions []Transaction
}
