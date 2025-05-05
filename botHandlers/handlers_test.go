package botHandlers

import (
	"testing"
	"unicode/utf8"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Positive cases
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "sentence with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "palindrome",
			input:    "radar",
			expected: "radar",
		},
		{
			name:     "unicode characters",
			input:    "привет 👋",
			expected: "👋 тевирп",
		},
		{
			name:     "numbers and special characters",
			input:    "123!@#",
			expected: "#@!321",
		},
		// Edge cases and potentially problematic inputs
		{
			name:     "very long string",
			input:    "Hello, this is a long string that we will reverse!",
			expected: "!esrever lliw ew taht gnirts gnol a si siht ,olleH",
		},
		{
			name:     "multiple spaces",
			input:    "   multiple   spaces   ",
			expected: "   secaps   elpitlum   ",
		},
		{
			name:     "only spaces",
			input:    "     ",
			expected: "     ",
		},
		{
			name:     "mixed unicode",
			input:    "Hello 世界 👋 Привет",
			expected: "тевирП 👋 界世 olleH",
		},
		{
			name:     "special unicode characters",
			input:    "⌘⌥⇧⌃",
			expected: "⌃⇧⌥⌘",
		},
		{
			name:     "zero-width characters",
			input:    "a\u200bb\u200bc", // contains zero-width space
			expected: "c\u200bb\u200ba",
		},
		{
			name:     "newlines and tabs",
			input:    "hello\nworld\ttab",
			expected: "bat\tdlrow\nolleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.input)
			if result != tt.expected {
				t.Errorf("reverseString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestReverseStringProperties tests properties that should always be true
func TestReverseStringProperties(t *testing.T) {
	tests := []struct {
		name     string
		property func(string) bool
		input    string
	}{
		{
			name: "double reverse returns original",
			property: func(s string) bool {
				return reverseString(reverseString(s)) == s
			},
			input: "Hello 世界 👋 Привет",
		},
		{
			name: "length remains the same",
			property: func(s string) bool {
				return utf8.RuneCountInString(reverseString(s)) == utf8.RuneCountInString(s)
			},
			input: "Hello 世界 👋 Привет",
		},
		{
			name: "spaces count remains the same",
			property: func(s string) bool {
				original := countSpaces(s)
				reversed := countSpaces(reverseString(s))
				return original == reversed
			},
			input: "  Hello  World  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.property(tt.input) {
				t.Errorf("Property %q failed for input %q", tt.name, tt.input)
			}
		})
	}
}

// Helper function to count spaces in a string
func countSpaces(s string) int {
	count := 0
	for _, r := range s {
		if r == ' ' {
			count++
		}
	}
	return count
}
