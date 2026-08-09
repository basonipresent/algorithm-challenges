package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example from doc comment",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "single character match",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "no valid window exists",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "target longer than source",
			s:    "a",
			t:    "aabbcc",
			want: "",
		},
		{
			name: "whole string is the minimum window",
			s:    "ab",
			t:    "ab",
			want: "ab",
		},
		{
			name: "repeated characters required",
			s:    "aa",
			t:    "aa",
			want: "aa",
		},
		{
			name: "target character not present in source at all",
			s:    "a",
			t:    "b",
			want: "",
		},
		{
			name: "tricky case with scattered duplicates",
			s:    "cabwefgewcwaefgcf",
			t:    "cae",
			want: "cwae",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
