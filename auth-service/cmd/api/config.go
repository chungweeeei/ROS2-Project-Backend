package main

import (
	"auth-service/data"
	"log"

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
