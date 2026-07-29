package main

import (
	"reflect"
	"testing"
)

func buildList(digits []int) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	for _, d := range digits {
		curr.Next = &ListNode{Val: d}
		curr = curr.Next
	}
	return dummyHead.Next
}

func listToSlice(l *ListNode) []int {
	var out []int
	for l != nil {
		out = append(out, l.Val)
		l = l.Next
	}
	return out
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "example: 342 + 465 = 807",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "both lists empty",
			l1:   []int{},
			l2:   []int{},
			want: nil,
		},
		{
			name: "carry propagates through all digits: 9999999 + 9999 = 10009998",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "different lengths without carry: 1 + 99 = 100",
			l1:   []int{1},
			l2:   []int{9, 9},
			want: []int{0, 0, 1},
		},
		{
			name: "one list empty: 0 + 5 = 5",
			l1:   []int{},
			l2:   []int{5},
			want: []int{5},
		},
		{
			name: "zero plus zero",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listToSlice(addTwoNumbers(buildList(tt.l1), buildList(tt.l2)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}
