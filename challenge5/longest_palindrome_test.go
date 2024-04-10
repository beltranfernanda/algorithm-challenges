package challenge5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},                         // Empty string
		{"a", "a"},                       // Single character string
		{"abcde", "a"},                   // Non-palindromic string
		{"abccba", "abccba"},             // Palindromic string with odd length
		{"abccbaa", "abccba"},            // Palindromic string with even length
		{"abccbaabccba", "abccbaabccba"}, // Palindromic string with even length
		{"racecar", "racecar"},           // Palindromic string with odd length
		{"apple", "pp"},                  // Longest palindrome in the string
		{"noon", "noon"},                 // Palindromic string with odd length
		{"deified", "deified"},           // Palindromic string with odd length
	}

	for _, tc := range testCases {
		result := LongestPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
		}
	}
}
