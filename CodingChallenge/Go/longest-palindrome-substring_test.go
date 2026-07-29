package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example: multiple valid odd-length palindromes",
			s:    "babad",
			want: "bab",
		},
		{
			name: "even-length palindrome",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "single character",
			s:    "a",
			want: "a",
		},
		{
			name: "entire string is a palindrome",
			s:    "racecar",
			want: "racecar",
		},
		{
			name: "no repeated characters, answer is any single char",
			s:    "abc",
			want: "a",
		},
		{
			name: "all same character",
			s:    "aaaa",
			want: "aaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			if len(got) != len(tt.want) {
				t.Errorf("longestPalindrome(%q) = %q (len %d), want len %d", tt.s, got, len(got), len(tt.want))
			}
		})
	}
}
