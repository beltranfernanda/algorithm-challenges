package challenge3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},    // "b"
		{"pwwkew", 3},   // "wke"
		{"", 0},         // empty string
		{" ", 1},        // single character
		{"dvdf", 3},     // "vdf"
		{"abcde", 5},    // "abcde"
		{"aabbcc", 2},   // "ab" or "bc"
		{"abcdefg", 7},  // "abcdefg"
		{"abcdefga", 7}, // "bcdefga"
		{"dvdf", 3},     // "vdf"
		{"tmmzuxt", 5},  // "mzuxt"
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := LengthOfLongestSubstring(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d for input %s", tc.expected, result, tc.input)
			}
		})
	}
}
