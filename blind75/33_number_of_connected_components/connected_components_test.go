package connected_components

import (
	"testing"
)

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "example 1 - two components",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			want:  2,
		},
		{
			name:  "example 2 - one component",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:  1,
		},
		{
			name:  "no edges - all isolated",
			n:     4,
			edges: [][]int{},
			want:  4,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  1,
		},
		{
			name:  "complete graph",
			n:     4,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			want:  1,
		},
		{
			name:  "three components",
			n:     6,
			edges: [][]int{{0, 1}, {2, 3}, {4, 5}},
			want:  3,
		},
		{
			name:  "star graph",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountComponents(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("CountComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountComponentsBFS(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "example 1 - two components",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			want:  2,
		},
		{
			name:  "example 2 - one component",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:  1,
		},
		{
			name:  "no edges",
			n:     4,
			edges: [][]int{},
			want:  4,
		},
		{
			name:  "complete graph",
			n:     4,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountComponentsBFS(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("CountComponentsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountComponentsUnionFind(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "example 1 - two components",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			want:  2,
		},
		{
			name:  "example 2 - one component",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:  1,
		},
		{
			name:  "no edges",
			n:     4,
			edges: [][]int{},
			want:  4,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  1,
		},
		{
			name:  "three components",
			n:     6,
			edges: [][]int{{0, 1}, {2, 3}, {4, 5}},
			want:  3,
		},
		{
			name:  "star graph",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountComponentsUnionFind(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("CountComponentsUnionFind() = %v, want %v", got, tt.want)
			}
		})
	}
}
