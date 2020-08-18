package utils_test

import (
	"testing"

	"github.com/lazhari/web-jwt/utils"
)

func TestIsEmailValid(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected bool
	}{
		{"invalid email", "_invalid@", false},
		{"invalid email domain", "test@xxxx_s.co", false},
		{"valid email", "test@golang.org", true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			isValid := utils.IsEmailValid(tc.input)

			if isValid != tc.expected {
				t.Errorf("Expect %v; got %v", tc.expected, isValid)
			}
		})
	}

}
