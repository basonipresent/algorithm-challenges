package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "single matching pair",
			s:    "()",
			want: true,
		},
		{
			name: "multiple distinct pairs in sequence",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "mismatched bracket types",
			s:    "(]",
			want: false,
		},
		{
			name: "improperly interleaved brackets",
			s:    "([)]",
			want: false,
		},
		{
			name: "properly nested brackets",
			s:    "{[]}",
			want: true,
		},
		{
			name: "empty string is valid",
			s:    "",
			want: true,
		},
		{
			name: "single unmatched opening bracket",
			s:    "(",
			want: false,
		},
		{
			name: "single unmatched closing bracket",
			s:    "]",
			want: false,
		},
		{
			name: "two unmatched opening brackets",
			s:    "((",
			want: false,
		},
		{
			name: "sequential pairs of different types",
			s:    "(){}[]",
			want: true,
		},
		{
			name: "closing bracket appears before its opener",
			s:    ")(",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
