package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "example from doc comment",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "no common prefix",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "single string",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "empty slice",
			strs: []string{},
			want: "",
		},
		{
			name: "one string is a prefix of another",
			strs: []string{"ab", "a"},
			want: "a",
		},
		{
			name: "all identical strings",
			strs: []string{"test", "test", "test"},
			want: "test",
		},
		{
			name: "shortest string first",
			strs: []string{"a", "ab", "abc"},
			want: "a",
		},
		{
			name: "contains an empty string",
			strs: []string{"abc", "", "abd"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
