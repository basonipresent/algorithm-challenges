package main

import (
	"reflect"
	"testing"
)

func TestTwoSumIISorted(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example from doc comment",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "pair at the ends",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "negative and positive numbers",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "duplicate values summing to target",
			numbers: []int{3, 3},
			target:  6,
			want:    []int{1, 2},
		},
		{
			name:    "adjacent pair in the middle",
			numbers: []int{1, 2, 3, 4, 4, 9, 56, 90},
			target:  8,
			want:    []int{4, 5},
		},
		{
			name:    "no pair sums to target",
			numbers: []int{1, 2, 3},
			target:  100,
			want:    nil,
		},
		{
			name:    "all negative numbers",
			numbers: []int{-5, -3, -1},
			target:  -8,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumIISorted(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumIISorted(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
