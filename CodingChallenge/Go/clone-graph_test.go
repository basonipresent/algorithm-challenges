package main

import (
	"reflect"
	"sort"
	"testing"
)

// buildGraph constructs a graph from an adjacency list where node values are
// 1-indexed (adjacency[i] holds the neighbor values for node i+1), matching
// the format used in the doc comment's example.
func buildGraph(adjacency [][]int) *Node {
	if len(adjacency) == 0 {
		return nil
	}
	nodes := make(map[int]*Node)
	for i := range adjacency {
		val := i + 1
		nodes[val] = &Node{Val: val}
	}
	for i, neighbors := range adjacency {
		val := i + 1
		for _, n := range neighbors {
			nodes[val].Neighbors = append(nodes[val].Neighbors, nodes[n])
		}
	}
	return nodes[1]
}

// graphAdjacency walks a graph via BFS and returns a map of node value to
// sorted neighbor values, for structural comparison independent of pointer
// identity or traversal/neighbor order.
func graphAdjacency(start *Node) map[int][]int {
	result := make(map[int][]int)
	if start == nil {
		return result
	}
	visited := make(map[*Node]bool)
	queue := []*Node{start}
	visited[start] = true
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		neighborVals := make([]int, 0, len(n.Neighbors))
		for _, nb := range n.Neighbors {
			neighborVals = append(neighborVals, nb.Val)
			if !visited[nb] {
				visited[nb] = true
				queue = append(queue, nb)
			}
		}
		sort.Ints(neighborVals)
		result[n.Val] = neighborVals
	}
	return result
}

// collectNodes returns the set of all node pointers reachable from start.
func collectNodes(start *Node) map[*Node]bool {
	visited := make(map[*Node]bool)
	if start == nil {
		return visited
	}
	queue := []*Node{start}
	visited[start] = true
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, nb := range n.Neighbors {
			if !visited[nb] {
				visited[nb] = true
				queue = append(queue, nb)
			}
		}
	}
	return visited
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name       string
		adjacency  [][]int
		wantVals   map[int][]int
		wantNilOut bool
	}{
		{
			name:      "example from doc comment",
			adjacency: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			wantVals:  map[int][]int{1: {2, 4}, 2: {1, 3}, 3: {2, 4}, 4: {1, 3}},
		},
		{
			name:      "single node, no neighbors",
			adjacency: [][]int{{}},
			wantVals:  map[int][]int{1: {}},
		},
		{
			name:       "nil input graph",
			adjacency:  [][]int{},
			wantNilOut: true,
		},
		{
			name:      "two nodes connected to each other",
			adjacency: [][]int{{2}, {1}},
			wantVals:  map[int][]int{1: {2}, 2: {1}},
		},
		{
			name:      "triangle of three nodes",
			adjacency: [][]int{{2, 3}, {1, 3}, {1, 2}},
			wantVals:  map[int][]int{1: {2, 3}, 2: {1, 3}, 3: {1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjacency)
			clone := cloneGraph(original)

			if tt.wantNilOut {
				if clone != nil {
					t.Fatalf("cloneGraph(nil) = %v, want nil", clone)
				}
				return
			}

			gotVals := graphAdjacency(clone)
			if !reflect.DeepEqual(gotVals, tt.wantVals) {
				t.Errorf("cloneGraph adjacency = %v, want %v", gotVals, tt.wantVals)
			}

			originalNodes := collectNodes(original)
			cloneNodes := collectNodes(clone)
			for n := range cloneNodes {
				if originalNodes[n] {
					t.Errorf("cloneGraph reused a pointer from the original graph: %v", n)
				}
			}
		})
	}
}
