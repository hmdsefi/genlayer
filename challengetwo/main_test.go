package main

import (
	"testing"
)

func TestMinimizeMaxRisk(t *testing.T) {
	tests := []struct {
		name        string
		dataCenters []int
		fragments   int
		expected    int
	}{
		{
			name:        "Example 1",
			dataCenters: []int{10, 20, 30},
			fragments:   2,
			expected:    1400, // Expected minimized maximum risk
		},
		{
			name:        "Example 2",
			dataCenters: []int{5, 10, 15},
			fragments:   3,
			expected:    4500, // Expected minimized maximum risk
		},
		{
			name:        "Single Data Center",
			dataCenters: []int{5},
			fragments:   4,
			expected:    625, // Expected minimized maximum risk
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actual := minimizeMaxRisk(tt.dataCenters, tt.fragments)
				if actual != tt.expected {
					t.Errorf("Test %s failed: expected %d, got %d", tt.name, tt.expected, actual)
				}
			},
		)
	}
}
