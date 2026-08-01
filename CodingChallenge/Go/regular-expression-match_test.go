package main

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "no metacharacters, exact match",
			s:    "abc",
			p:    "abc",
			want: true,
		},
		{
			name: "no metacharacters, mismatch",
			s:    "abc",
			p:    "abd",
			want: false,
		},
		{
			name: "single char pattern shorter than string",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "star matches zero or more of preceding char",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "dot-star matches any string",
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			name: "star allows pattern shorter than matched string",
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			name: "classic mismatch example",
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
		{
			name: "star matching a run of the same character",
			s:    "aaa",
			p:    "a*a",
			want: true,
		},
		{
			name: "both empty",
			s:    "",
			p:    "",
			want: true,
		},
		{
			name: "empty string with star pattern matching zero occurrences",
			s:    "",
			p:    "a*",
			want: true,
		},
		{
			name: "dot matches any single character",
			s:    "abc",
			p:    "a.c",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
