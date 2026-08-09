package main

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "example from doc comment",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "alternating characters, enough replacements for whole string",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "already all same character, no replacement needed",
			s:    "AAAA",
			k:    0,
			want: 4,
		},
		{
			name: "single character",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "one minority character within budget",
			s:    "ABBB",
			k:    2,
			want: 4,
		},
		{
			name: "extra unused budget",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "empty string",
			s:    "",
			k:    0,
			want: 0,
		},
		{
			name: "zero budget with mixed characters caps window at one",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
