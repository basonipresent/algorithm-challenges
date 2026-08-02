package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example from doc comment",
			s:    "MCMXCIV",
			want: 1994,
		},
		{
			name: "simple additive symbols",
			s:    "III",
			want: 3,
		},
		{
			name: "subtractive pair for four",
			s:    "IV",
			want: 4,
		},
		{
			name: "subtractive pair for nine",
			s:    "IX",
			want: 9,
		},
		{
			name: "mix of additive and subtractive",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "single symbol",
			s:    "M",
			want: 1000,
		},
		{
			name: "subtractive pair for forty",
			s:    "XL",
			want: 40,
		},
		{
			name: "subtractive pair for ninety",
			s:    "XC",
			want: 90,
		},
		{
			name: "subtractive pair for four hundred",
			s:    "CD",
			want: 400,
		},
		{
			name: "subtractive pair for nine hundred",
			s:    "CM",
			want: 900,
		},
		{
			name: "max value for this problem",
			s:    "MMMCMXCIX",
			want: 3999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.want {
				t.Errorf("romanToInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
