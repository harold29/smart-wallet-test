package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbInfo string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
