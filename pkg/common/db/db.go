package db

import (
	"harold29/yourkeyswallet/pkg/common/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(dbInfo string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserType{},
		&models.UserRole{},
		&models.Wallet{},
		&models.Transaction{},
		&models.Currency{},
		&models.TransactionType{},
		&models.Location{},
		&models.Geolocation{},
	)

	CreateSupportObjects(db)

	DB = db

	return DB
}

func ClearTable() {
	DB.Exec("DELETE FROM users")
	DB.Exec("DELETE FROM user_types")
	DB.Exec("DELETE FROM user_roles")
}

func CreateSupportObjects(db *gorm.DB) {
	var UserTypes = []models.UserType{{Name: "basic", Description: ""}, {Name: "advanced", Description: ""}}

	var UserRoles = []models.UserRole{{Name: "basic", Description: ""}, {Name: "normal", Description: ""}}

	db.Create(&UserTypes)
	db.Create(&UserRoles)
}
