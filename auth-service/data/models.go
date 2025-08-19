package data

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func New(dbPool *gorm.DB) Models {
	db = dbPool

	// Do auto migration
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println("Failed to auto migrate user table")
	}

	return Models{
		User: &User{},
	}
}

type Models struct {
	User UserInterface
}
