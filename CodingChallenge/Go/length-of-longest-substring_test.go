package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example with repeat later in string",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "all same character",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "repeat immediately followed by longer unique run",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "all unique characters",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "duplicate at the very start",
			s:    "aab",
			want: 2,
		},
		{
			name: "duplicate index before current window (stale map entry)",
			s:    "abba",
			want: 2,
		},
		{
			name: "string with spaces",
			s:    "a ab",
			want: 3,
		},
		{
			name: "unicode/multi-byte runes",
			s:    "日本語日本",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
