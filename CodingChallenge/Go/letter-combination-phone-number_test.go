package main

import (
	"sort"
	"testing"
)

func sortedCopy(s []string) []string {
	out := make([]string, len(s))
	copy(out, s)
	sort.Strings(out)
	return out
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "example from doc comment",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "empty input",
			digits: "",
			want:   []string{},
		},
		{
			name:   "single digit",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "digit with four letters",
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			name:   "three digits",
			digits: "234",
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			gotSorted := sortedCopy(got)
			wantSorted := sortedCopy(tt.want)
			if len(gotSorted) != len(wantSorted) {
				t.Fatalf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
			}
			for i := range gotSorted {
				if gotSorted[i] != wantSorted[i] {
					t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
					break
				}
			}
		})
	}
}
