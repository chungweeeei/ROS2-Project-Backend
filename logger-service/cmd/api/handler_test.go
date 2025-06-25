package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadAllLogs(t *testing.T) {

	router := testApp.routes()

	// Create a new HTTP request to the /v1/logs endpoint
	req, _ := http.NewRequest("GET", "/v1/logs", nil)
	rr := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check if the status code is OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, resp["error"], false)

	logs := resp["logs"].([]interface{})
	assert.Len(t, logs, 2, "Expected 2 logs in response")

	firstLogEntry := logs[0].(map[string]interface{})
	assert.Equal(t, firstLogEntry["name"], "Test Log")
	assert.Equal(t, firstLogEntry["level"], "info")
	assert.Equal(t, firstLogEntry["message"], "This is the first test log message")

	secondLogEntry := logs[1].(map[string]interface{})
	assert.Equal(t, secondLogEntry["name"], "Test Log")
	assert.Equal(t, secondLogEntry["level"], "info")
	assert.Equal(t, secondLogEntry["message"], "This is the second test log message")

	fmt.Printf("\033[32m✅ Finished Test Read All Logs Handler\033[0m\n")
}
