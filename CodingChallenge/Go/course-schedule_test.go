package main

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example from doc comment",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "direct cycle between two courses",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "no prerequisites at all",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "longer cycle across three courses",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			want:          false,
		},
		{
			name:          "diamond dependency, no cycle",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "single course, no prerequisites",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "disconnected components, one with a cycle",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {3, 2}, {2, 3}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
