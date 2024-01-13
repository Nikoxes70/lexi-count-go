package wordsbank

import (
	"testing"
)

func TestIsValidWord(t *testing.T) {
	v := NewValidator()

	testCases := []struct {
		word string
		want bool
	}{
		{"example", true},
		{"hi", false},        // Less than 3 characters
		{"1234", false},      // Non-alphabetic characters
		{"hello2", false},    // Mixed characters
		{"Hello", true},      // Valid with uppercase
		{"", false},          // Empty string
		{"  space  ", false}, // Contains spaces
		{"two words", false}, // More than one word
		{"£lan", false},      // Non-ASCII characters
		{"abc123", false},    // Alphanumeric mix
		{"One", true},        // Valid with first character uppercase
		{"UPPERCASE", true},  // All uppercase letters
		{"lowercase", true},  // All lowercase letters
		{"Special$", false},  // Special character inclusion
		{"newline\n", false}, // Newline character
		{"tab\tchar", false}, // Tab character
	}

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			got := v.IsValidWord(tc.word)
			if got != tc.want {
				t.Errorf("IsValidWord(%q) = %v, want %v", tc.word, got, tc.want)
			}
		})
	}
}
