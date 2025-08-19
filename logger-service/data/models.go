package data

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func New(dbPool *gorm.DB) Models {
	db = dbPool

	// Do auto migration
	err := db.AutoMigrate(&Log{})
	if err != nil {
		log.Println("Failed to auto migrate log table")
	}

	return Models{
		Log: &Log{},
	}
}

type Models struct {
	Log LogInterface
}
