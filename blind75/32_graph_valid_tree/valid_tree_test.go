package valid_tree

import (
	"testing"
)

func TestValidTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{
			name:  "example 1 - valid tree",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			want:  true,
		},
		{
			name:  "example 2 - has cycle",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			want:  false,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  true,
		},
		{
			name:  "two nodes connected",
			n:     2,
			edges: [][]int{{0, 1}},
			want:  true,
		},
		{
			name:  "disconnected graph",
			n:     4,
			edges: [][]int{{0, 1}, {2, 3}},
			want:  false,
		},
		{
			name:  "too many edges",
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:  false,
		},
		{
			name:  "too few edges",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}},
			want:  false,
		},
		{
			name:  "linear chain",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidTree(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("ValidTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidTreeUnionFind(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{
			name:  "example 1 - valid tree",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			want:  true,
		},
		{
			name:  "example 2 - has cycle",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			want:  false,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  true,
		},
		{
			name:  "two nodes connected",
			n:     2,
			edges: [][]int{{0, 1}},
			want:  true,
		},
		{
			name:  "disconnected graph",
			n:     4,
			edges: [][]int{{0, 1}, {2, 3}},
			want:  false,
		},
		{
			name:  "too many edges",
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:  false,
		},
		{
			name:  "linear chain",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidTreeUnionFind(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("ValidTreeUnionFind() = %v, want %v", got, tt.want)
			}
		})
	}
}
