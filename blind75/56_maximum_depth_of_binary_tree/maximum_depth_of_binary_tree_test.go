package maximum_depth_of_binary_tree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example 1 - balanced tree",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "example 2 - right skewed",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: 2,
		},
		{
			name: "example 3 - empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "example 4 - single node",
			root: &TreeNode{Val: 0},
			want: 1,
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			want: 4,
		},
		{
			name: "complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.root)
			if got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
