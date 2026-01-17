package binary_tree_maximum_path_sum

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example 1 - simple tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: 6,
		},
		{
			name: "example 2 - tree with negative root",
			root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 42,
		},
		{
			name: "example 3 - single negative node",
			root: &TreeNode{Val: -3},
			want: -3,
		},
		{
			name: "single positive node",
			root: &TreeNode{Val: 5},
			want: 5,
		},
		{
			name: "all negative values",
			root: &TreeNode{
				Val:   -2,
				Left:  &TreeNode{Val: -1},
				Right: &TreeNode{Val: -3},
			},
			want: -1,
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 3},
				},
			},
			want: 12,
		},
		{
			name: "tree with negative branch",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   -1,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: -2},
			},
			want: 6,
		},
		{
			name: "complex tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 11},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 13},
					Right: &TreeNode{Val: 4},
				},
			},
			want: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxPathSum(tt.root)
			if got != tt.want {
				t.Errorf("MaxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
