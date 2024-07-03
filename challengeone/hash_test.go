package main

import (
	"testing"
)

func TestSimpleHash(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"example input string", 30},
		{"another example", 30},
		{"short", 30},
		{"this is a longer input string to see how it handles more data", 30},
	}

	hash := NewHash()
	for _, tt := range tests {
		t.Run(
			tt.input, func(t *testing.T) {
				hash := hash.SimpleHash(tt.input)
				if len(hash) != tt.expected {
					t.Errorf("SimpleHash(%q) = %d characters; want %d", tt.input, len(hash), tt.expected)
				}
			},
		)
	}
}

func TestSimpleHashUniqueness(t *testing.T) {
	hash := NewHash()
	hash1 := hash.SimpleHash("input1")
	hash2 := hash.SimpleHash("input2")
	if hash1 == hash2 {
		t.Errorf("Expected different hashes for different inputs, but got the same: %s", hash1)
	}
}

func TestSimpleHashConsistency(t *testing.T) {
	input := "consistent input"

	hash := NewHash()
	hash1 := hash.SimpleHash(input)
	hash2 := hash.SimpleHash(input)
	if hash1 != hash2 {
		t.Errorf("Expected same hash for same input, but got different hashes: %s and %s", hash1, hash2)
	}
}
