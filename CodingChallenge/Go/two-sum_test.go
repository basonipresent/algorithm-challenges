package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example pair in middle",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "pair at the end",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "duplicate values summing to target",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "no solution exists",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   []int{},
		},
		{
			name:   "empty input",
			nums:   []int{},
			target: 0,
			want:   []int{},
		},
		{
			name:   "single element cannot pair with itself",
			nums:   []int{5},
			target: 10,
			want:   []int{},
		},
		{
			name:   "zeroes summing to target zero",
			nums:   []int{0, 4, 0},
			target: 0,
			want:   []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
