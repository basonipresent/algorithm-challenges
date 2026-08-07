package main

import (
	"sort"
	"strings"
	"testing"
)

// groupKeys normalizes a group-anagrams result into a sorted list of
// comma-joined, alphabetically-sorted groups so tests can compare results
// regardless of group/word ordering.
func groupKeys(groups [][]string) []string {
	keys := make([]string, len(groups))
	for i, g := range groups {
		sorted := append([]string(nil), g...)
		sort.Strings(sorted)
		keys[i] = strings.Join(sorted, ",")
	}
	sort.Strings(keys)
	return keys
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example from doc comment",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name: "single empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "single word",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "no anagrams among words",
			strs: []string{"abc", "def"},
			want: [][]string{{"abc"}, {"def"}},
		},
		{
			name: "repeated identical word",
			strs: []string{"ab", "ab", "ab"},
			want: [][]string{{"ab", "ab", "ab"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "multiple empty strings group together",
			strs: []string{"", ""},
			want: [][]string{{"", ""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			gotKeys := groupKeys(got)
			wantKeys := groupKeys(tt.want)
			if len(gotKeys) != len(wantKeys) {
				t.Fatalf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
			for i := range gotKeys {
				if gotKeys[i] != wantKeys[i] {
					t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
					break
				}
			}
		})
	}
}
