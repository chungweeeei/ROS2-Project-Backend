package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertStringToTime(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "valid date",
			input:    "2025-08-20",
			expected: time.Date(2025, 8, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "invalid date format",
			input:    "20-08-2025",
			expected: time.Time{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := convertStringToTime(tc.input)
			assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected %v but got %v for input %s", tc.expected, result, tc.input))
		})
	}

}
