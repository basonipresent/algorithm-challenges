package main

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "example from doc comment",
			num:  3749,
			want: "MMMDCCXLIX",
		},
		{
			name: "single symbol",
			num:  1,
			want: "I",
		},
		{
			name: "subtractive pair for four",
			num:  4,
			want: "IV",
		},
		{
			name: "subtractive pair for nine",
			num:  9,
			want: "IX",
		},
		{
			name: "classic 1994",
			num:  1994,
			want: "MCMXCIV",
		},
		{
			name: "classic 58",
			num:  58,
			want: "LVIII",
		},
		{
			name: "max value for this problem",
			num:  3999,
			want: "MMMCMXCIX",
		},
		{
			name: "multiple of a repeated symbol",
			num:  3000,
			want: "MMM",
		},
		{
			name: "forty uses subtractive notation",
			num:  40,
			want: "XL",
		},
		{
			name: "ninety uses subtractive notation",
			num:  90,
			want: "XC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Errorf("intToRoman(%d) = %q, want %q", tt.num, got, tt.want)
			}
		})
	}
}
