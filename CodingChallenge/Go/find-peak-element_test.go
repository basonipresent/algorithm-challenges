package main

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example from doc comment",
			nums: []int{1, 2, 3, 1},
			want: 2,
		},
		{
			name: "multiple peaks, any valid answer accepted",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			want: 5,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 0,
		},
		{
			name: "peak at the start",
			nums: []int{5, 4, 3, 2, 1},
			want: 0,
		},
		{
			name: "peak at the end",
			nums: []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			name: "two elements, first is peak",
			nums: []int{2, 1},
			want: 0,
		},
		{
			name: "two elements, second is peak",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "peak in the middle of a valley shape",
			nums: []int{1, 3, 2, 4, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.nums)
			if got < 0 || got >= len(tt.nums) {
				t.Fatalf("findPeakElement(%v) = %d, out of bounds", tt.nums, got)
			}
			if tt.want != got {
				t.Errorf("findPeakElement(%v) = %d, but want %d", tt.nums, got, tt.want)
			}
		})
	}
}
