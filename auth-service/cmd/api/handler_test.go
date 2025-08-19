package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Authenticate(t *testing.T) {

	// Get the router
	router := testApp.routes()

	// Define test cases
	tests := []struct {
		name               string
		payload            map[string]string
		expectedStatusCode int
		expectedResponse   map[string]interface{}
	}{
		{
			name: "valid login",
			payload: map[string]string{
				"email":    "admin@example.com",
				"password": "abc",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: map[string]interface{}{
				"email":   "admin@example.com",
				"message": "Authenticate successfully",
			},
		},
		{
			name: "wrong request",
			payload: map[string]string{
				"password": "abc",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"message": "Failed to parse request body",
			},
		},
		{
			name: "wrong user",
			payload: map[string]string{
				"email":    "test-wrong@example.com",
				"password": "abc",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{
				"message": "Invalid email or password",
			},
		}, {
			name: "wrong password",
			payload: map[string]string{
				"email":    "admin@example.com",
				"password": "wrong",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{
				"message": "Invalid email or password",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error

			jsonData, _ := json.Marshal(tc.payload)
			req, err = http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")

			// Create a Response Recorder
			rr := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(rr, req)

			// Check the status code
			if rr.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, rr.Code)
			}

			// Check the response body
			var resp map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &resp)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for key, expectedValue := range tc.expectedResponse {
				assert.Equal(t, expectedValue, resp[key])
			}
		})
	}

	testApp.InfoLog.Printf("\033[32m✅ Finished Test Login Handler\033[0m\n")
}

func Test_Signup(t *testing.T) {

	// Get the router
	router := testApp.routes()

	// Define test cases
	test := []struct {
		name               string
		payload            map[string]string
		expectedStatusCode int
		expectedResponse   map[string]interface{}
	}{
		{
			name: "valid signup",
			payload: map[string]string{
				"email":      "test@test.com",
				"first_name": "Andy",
				"last_name":  "Tseng",
				"password":   "12345678",
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: map[string]interface{}{
				"message": "User created successfully",
			},
		},
		{
			name: "failed insert",
			payload: map[string]string{
				"email":      "test-wrong@example.com",
				"first_name": "Andy",
				"last_name":  "Tseng",
				"password":   "12345678",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse: map[string]interface{}{
				"message": "Failed to register new user",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error

			jsonData, _ := json.Marshal(tc.payload)
			req, err = http.NewRequest("POST", "/v1/signup", bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")

			// Create a Response Recorder
			rr := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(rr, req)

			// Check the status code
			if rr.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, rr.Code)
			}

			// Check the response body
			var resp map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &resp)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for key, expectedValue := range tc.expectedResponse {
				assert.Equal(t, expectedValue, resp[key])
			}
		})
	}
}
