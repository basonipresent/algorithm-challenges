package main

import "testing"

// buildListWithCycle builds a linked list from vals and, if pos >= 0,
// makes the tail point back to the node at index pos to form a cycle.
func buildListWithCycle(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{
			name: "example from doc comment",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			name: "two-node cycle",
			vals: []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			name: "single node self cycle",
			vals: []int{1},
			pos:  0,
			want: true,
		},
		{
			name: "cycle near end of longer list",
			vals: []int{1, 2, 3, 4, 5},
			pos:  3,
			want: true,
		},
		{
			name: "single node no cycle",
			vals: []int{1},
			pos:  -1,
			want: false,
		},
		{
			name: "two nodes no cycle",
			vals: []int{1, 2},
			pos:  -1,
			want: false,
		},
		{
			name: "multiple nodes no cycle",
			vals: []int{1, 2, 3, 4},
			pos:  -1,
			want: false,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListWithCycle(tt.vals, tt.pos)
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle(%v, pos=%d) = %v, want %v", tt.vals, tt.pos, got, tt.want)
			}
		})
	}
}
