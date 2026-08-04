package main

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{
			name: "remove from middle",
			nums: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "remove only node",
			nums: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "remove head of two-node list",
			nums: []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "remove tail",
			nums: []int{1, 2, 3},
			n:    1,
			want: []int{1, 2},
		},
		{
			name: "remove head of longer list",
			nums: []int{1, 2, 3, 4, 5},
			n:    5,
			want: []int{2, 3, 4, 5},
		},
		{
			name: "remove head of middle list",
			nums: []int{1, 2, 3, 4, 5},
			n:    4,
			want: []int{1, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listToSlice(removeNthFromEnd(buildList(tt.nums), tt.n))
			if !reflect.DeepEqual(got, tt.want) && !(len(got) == 0 && len(tt.want) == 0) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", tt.nums, tt.n, got, tt.want)
			}
		})
	}
}
