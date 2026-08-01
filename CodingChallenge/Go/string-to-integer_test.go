package main

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "simple positive number",
			str:  "42",
			want: 42,
		},
		{
			name: "leading spaces with negative sign",
			str:  "   -42",
			want: -42,
		},
		{
			name: "stops at first non-digit character",
			str:  "4193 with words",
			want: 4193,
		},
		{
			name: "no leading digits returns zero",
			str:  "words and 987",
			want: 0,
		},
		{
			name: "negative overflow clamps to int32 min",
			str:  "-91283472332",
			want: -2147483648,
		},
		{
			name: "positive overflow clamps to int32 max",
			str:  "91283472332",
			want: 2147483647,
		},
		{
			name: "empty string",
			str:  "",
			want: 0,
		},
		{
			name: "only spaces",
			str:  "   ",
			want: 0,
		},
		{
			name: "explicit plus sign",
			str:  "+1",
			want: 1,
		},
		{
			name: "stops at space after sign and digits",
			str:  "  +0 123",
			want: 0,
		},
		{
			name: "stops at decimal point",
			str:  "3.14159",
			want: 3,
		},
		{
			name: "digits then letter then more digits",
			str:  "   -0012a42",
			want: -12,
		},
		{
			name: "one more than int32 max",
			str:  "2147483648",
			want: 2147483647,
		},
		{
			name: "exactly int32 min",
			str:  "-2147483648",
			want: -2147483648,
		},
		{
			name: "leading zeros",
			str:  "  0000000000012345678",
			want: 12345678,
		},
		{
			name: "sign with no digits",
			str:  "+",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.str)
			if got != tt.want {
				t.Errorf("atoi(%q) = %d, want %d", tt.str, got, tt.want)
			}
		})
	}
}
