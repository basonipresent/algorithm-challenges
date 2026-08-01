package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "positive palindrome, odd digit count",
			x:    121,
			want: true,
		},
		{
			name: "not a palindrome",
			x:    123,
			want: false,
		},
		{
			name: "negative number is never a palindrome",
			x:    -121,
			want: false,
		},
		{
			name: "zero is a palindrome",
			x:    0,
			want: true,
		},
		{
			name: "single digit is always a palindrome",
			x:    7,
			want: true,
		},
		{
			name: "palindrome with even digit count",
			x:    1221,
			want: true,
		},
		{
			name: "trailing zero, not a palindrome",
			x:    10,
			want: false,
		},
		{
			name: "large palindrome",
			x:    1234321,
			want: true,
		},
		{
			name: "large non-palindrome",
			x:    1234322,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
