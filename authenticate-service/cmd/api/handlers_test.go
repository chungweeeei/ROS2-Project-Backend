package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a Mock http client that allows us to mock the HTTP requests and responses
type RoundTripFunc func(req *http.Request) *http.Response

// It a placeholder function
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		// override the default transport with our custom function
		// we can substitute this with a mock or a stub depending on the test case
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {

	router := testApp.routes()

	// mock the response
	jsonToReturn := `
{
	"error": false,
	"message": "some message"
}
	`
	// mock the HTTP client to return the mocked response
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

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

// func Test_Signup(t *testing.T) {

// 	router := testApp.routes()

// 	// create test user request
// 	postBody := map[string]interface{}{
// 		"email":    "test@test.com",
// 		"username": "testuser",
// 		"password": "tester",
// 	}

// 	body, _ := json.Marshal(postBody)

// 	req, _ := http.NewRequest("POST", "/v1/authenticate/signup", bytes.NewBuffer(body))
// 	rr := httptest.NewRecorder()

// 	router.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusCreated {
// 		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
// 	}

// 	var resp map[string]interface{}
// 	err := json.Unmarshal(rr.Body.Bytes(), &resp)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal response: %v", err)
// 	}

// 	assert.Equal(t, resp["error"], false)
// 	assert.Equal(t, resp["message"], "User created successfully")

// 	fmt.Printf("\033[32m✅ Finished Test Signup Handler\033[0m\n")
// }
