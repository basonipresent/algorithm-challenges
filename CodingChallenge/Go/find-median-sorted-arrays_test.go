package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "example: combined odd length",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "example: combined even length",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "first array empty",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  2.0,
		},
		{
			name:  "second array empty",
			nums1: []int{2},
			nums2: []int{},
			want:  2.0,
		},
		{
			name:  "no overlap, nums1 entirely smaller",
			nums1: []int{1, 2},
			nums2: []int{3, 4, 5},
			want:  3.0,
		},
		{
			name:  "no overlap, nums2 entirely smaller",
			nums1: []int{4, 5, 6},
			nums2: []int{1, 2, 3},
			want:  3.5,
		},
		{
			name:  "single element each",
			nums1: []int{1},
			nums2: []int{2},
			want:  1.5,
		},
		{
			name:  "duplicate values across arrays",
			nums1: []int{2, 2},
			nums2: []int{2, 2},
			want:  2.0,
		},
		{
			name:  "negative numbers",
			nums1: []int{-5, -3, -1},
			nums2: []int{-4, -2},
			want:  -3.0,
		},
		{
			name:  "one array fully interleaved with the other",
			nums1: []int{1, 3, 5, 7},
			nums2: []int{2, 4, 6, 8},
			want:  4.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
