package main

import (
	"fmt"
	"logger-service/data"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testApp Config

func setup() {
	gin.SetMode(gin.TestMode)
	repo, _ := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	fmt.Println("Setting up the logger-service testing environment...")
}

func teardown() {
	fmt.Println("Tearing down the logger-service testing environment...")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run() // Run the tests
	teardown()
	os.Exit(code)
}
