package main

import (
	"log"
	"record-service/data"
	"record-service/gateways"

	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	AuthClient    gateways.AuthServiceClient
	Models        data.Models
	ErrorChan     chan error
	ErrorDoneChan chan bool
}
