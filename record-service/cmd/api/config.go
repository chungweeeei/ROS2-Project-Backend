package main

import (
	"log"
	"record-service/data"

	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Models        data.Models
	ErrorChan     chan error
	ErrorDoneChan chan bool
}
