package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Authenticate(t *testing.T) {

	router := testApp.routes()

	// create test user request
	postBody := map[string]interface{}{
		"email":    "test@test.com",
		"password": "tester",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/v1/authenticate/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, resp["error"], false)
	assert.Equal(t, resp["email"], "test@test.com")
	assert.Equal(t, resp["message"], "Authentication successful")
	assert.Contains(t, resp, "token")

	fmt.Printf("\033[32m✅ Finished Test Login Handler\033[0m\n")
}

func Test_Signup(t *testing.T) {

	router := testApp.routes()

	// create test user request
	postBody := map[string]interface{}{
		"email":    "test@test.com",
		"username": "testuser",
		"password": "tester",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/v1/authenticate/signup", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, resp["error"], false)
	assert.Equal(t, resp["message"], "User created successfully")

	fmt.Printf("\033[32m✅ Finished Test Signup Handler\033[0m\n")
}
