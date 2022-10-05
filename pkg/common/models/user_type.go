package models

import "gorm.io/gorm"

type UserType struct {
	gorm.Model
	ID          string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserType    string
	Description string
}
