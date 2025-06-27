package main

import (
	"authenticate-service/logs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
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

// Define a Mock LogServiceClient
type MockLogServiceClient struct {
	WriteLogFunc func(ctx context.Context, in *logs.LogRequest, opts ...grpc.CallOption) (*logs.LogResponse, error)
	CallCount    int
	LastRequest  *logs.LogRequest
}

func (m *MockLogServiceClient) WriteLog(ctx context.Context, req *logs.LogRequest, opts ...grpc.CallOption) (*logs.LogResponse, error) {
	m.CallCount++
	m.LastRequest = req

	if m.WriteLogFunc != nil {
		return m.WriteLogFunc(ctx, req)
	}

	// Default successful response
	return &logs.LogResponse{
		Result: "Log written successfully",
	}, nil
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

	testApp.Clients.LogHTTPClient = client

	gRPCClient := &MockLogServiceClient{
		WriteLogFunc: func(ctx context.Context, req *logs.LogRequest, opts ...grpc.CallOption) (*logs.LogResponse, error) {
			return &logs.LogResponse{
				Result: "Log written successfully",
			}, nil
		},
	}

	testApp.Clients.LoggRPCClient = gRPCClient

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

	fmt.Println("Response:", resp)

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
