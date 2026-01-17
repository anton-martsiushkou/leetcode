package course_schedule

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example 1 - possible",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "example 2 - impossible cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "no prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "linear chain",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "complex valid graph",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "cycle in middle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 1}},
			want:          false,
		},
		{
			name:          "self loop",
			numCourses:    1,
			prerequisites: [][]int{{0, 0}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("CanFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanFinishBFS(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example 1 - possible",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "example 2 - impossible cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "complex valid graph",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "cycle in middle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 1}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanFinishBFS(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("CanFinishBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
