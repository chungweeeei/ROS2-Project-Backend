package data

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func New(dbPool *gorm.DB) Models {

	db = dbPool

	err := db.AutoMigrate(&TradeRecord{})
	if err != nil {
		log.Println("Failed to auto migrate trade record table")
	}

	return Models{
		TradeRecord: &TradeRecord{},
	}

}

type Models struct {
	TradeRecord TradeRecordInterface
}
