package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
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
