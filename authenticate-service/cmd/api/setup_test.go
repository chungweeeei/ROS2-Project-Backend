package main

import (
	"authenticate-service/data"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testApp Config

func setup() {
	gin.SetMode(gin.TestMode)
	repo, _ := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	fmt.Println("Setting up the testing environment...")
}

func teardown() {
	fmt.Println("Tearing down the testing environment...")
}

// TestMain function is used to set up the environment before running tests.
func TestMain(m *testing.M) {

	setup()
	code := m.Run() // Run the tests
	teardown()
	os.Exit(code)
}
