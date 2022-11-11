package models

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID     `gorm:"primaryKey; unique; type:uuid; column:id; default:uuid_generate_v4()" json:"_"`
	FirstName    string        `gorm:"not null" json:"first_name"`
	LastName     string        `gorm:"not null" json:"last_name"`
	Email        string        `gorm:"not null; unique" json:"email"`
	PhoneNumber1 string        `json:"phone_number_1"`
	PhoneNumber2 string        `json:"phone_number_2"`
	Gender       string        `json:"gender"`
	Birthday     time.Time     `gorm:"not null" json:"birthday"`
	UserTypeID   uuid.UUID     `json:"user_type_uid"`
	UserType     UserType      `gorm:"foreignKey:UserTypeID" json:"user_type"`
	UserRoleID   uuid.UUID     `json:"user_role_uid"`
	UserRole     UserRole      `gorm:"foreignKey:UserRoleID" json:"user_role"`
	Wallets      []Wallet      `json:"wallets"`
	Transactions []Transaction `gorm:"many2many:user_transactions;" json:"transactions"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var role UserRole
	var uType UserType

	err1 := tx.First(&role, "name = ?", "basic").Error
	err2 := tx.First(&uType, "name = ?", "basic").Error

	if err1 != nil && err2 != nil {
		err = errors.New("can't save invalid data")
	}

	u.UserRole = role
	u.UserRoleID = role.ID
	u.UserType = uType
	u.UserTypeID = uType.ID

	return
}
