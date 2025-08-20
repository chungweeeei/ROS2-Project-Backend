package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchAllRecords(t *testing.T) {

	// Get the router
	router := testApp.routes()

	// Define test cases
	tests := []struct {
		name               string
		token              string
		payload            map[string]string
		expectedStatusCode int
		expectedResponse   map[string]interface{}
	}{
		{
			name:               "invalid token",
			token:              "",
			payload:            map[string]string{},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{
				"message": "authorization header not provided",
			},
		},
		{
			name:               "invalid token",
			token:              "Bearer",
			payload:            map[string]string{},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{
				"message": "invalid authorization header format",
			},
		},
		{
			name:               "valid request",
			token:              "Bearer valid-token",
			payload:            map[string]string{},
			expectedStatusCode: http.StatusOK,
			expectedResponse: map[string]interface{}{
				"records": []interface{}{
					map[string]interface{}{
						"email":        "test@test.com",
						"stock_number": "2330",
						"stock_name":   "台積電",
						"side":         "buy",
						"entry_price":  1185.0,
						"exit_price":   1200.0,
						"quantity":     1000.0,
						"entry_time":   "2025-08-18T00:00:00Z",
						"exit_time":    "2025-08-19T00:00:00Z",
						"notes":        "Test record",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error

			req, err = http.NewRequest("GET", "/v1/records", nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			req.Header.Set("Authorization", tc.token)

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

			if rr.Code == http.StatusUnauthorized {
				assert.Equal(t, tc.expectedResponse, resp, "Expected error message to match")
				return
			}

			assert.Len(t, resp["records"], 1, "Expected one record in response")

			record := resp["records"].([]interface{})[0].(map[string]interface{})
			expectedRecord := tc.expectedResponse["records"].([]interface{})[0].(map[string]interface{})

			assert.Equal(t, expectedRecord["email"], record["email"], "Expected email to match")
			assert.Equal(t, expectedRecord["stock_number"], record["stock_number"], "Expected stock number to match")
			assert.Equal(t, expectedRecord["stock_name"], record["stock_name"], "Expected stock name to match")
			assert.Equal(t, expectedRecord["side"], record["side"], "Expected side to match")
			assert.Equal(t, expectedRecord["entry_price"], record["entry_price"], "Expected entry price to match")
			assert.Equal(t, expectedRecord["exit_price"], record["exit_price"], "Expected exit price to match")
			assert.Equal(t, expectedRecord["quantity"], record["quantity"], "Expected quantity to match")
			assert.Equal(t, expectedRecord["entry_time"], record["entry_time"], "Expected entry time to match")
			assert.Equal(t, expectedRecord["exit_time"], record["exit_time"], "Expected exit time to match")
		})
	}
}
