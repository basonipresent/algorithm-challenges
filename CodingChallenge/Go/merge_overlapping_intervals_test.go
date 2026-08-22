package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "example from doc comment",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}},
			want:      [][]int{{1, 2}, {3, 4}},
		},
		{
			name:      "touching intervals merge",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "chain of overlaps merges into one",
			intervals: [][]int{{1, 4}, {2, 5}, {3, 6}},
			want:      [][]int{{1, 6}},
		},
		{
			name:      "interval fully nested inside another",
			intervals: [][]int{{1, 10}, {2, 5}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "unsorted input still merges correctly",
			intervals: [][]int{{5, 6}, {1, 3}, {2, 4}},
			want:      [][]int{{1, 4}, {5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
