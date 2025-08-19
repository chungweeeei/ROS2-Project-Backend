package main

import (
	"auth-service/data"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testApp Config

func setup() {
	gin.SetMode(gin.TestMode)
	testApp = Config{
		DB:            nil,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Models:        data.TestNew(nil),
		ErrorChan:     make(chan error),
		ErrorDoneChan: make(chan bool),
	}
	testApp.InfoLog.Println("Setting up the auth service testing environment...")
}

func teardown() {
	testApp.InfoLog.Println("Tearing down the auth service testing environment...")
}

func TestMain(m *testing.M) {

	setup()
	code := m.Run()
	teardown()
	os.Exit(code)

}
