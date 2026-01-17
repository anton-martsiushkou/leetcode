package invert_binary_tree

import (
	"testing"
)

// Helper function to check if two trees are identical
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "example 1 - complete tree",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "example 2 - simple tree",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name: "example 3 - empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: &TreeNode{Val: 1},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertTree(tt.root)
			if !isSameTree(got, tt.want) {
				t.Errorf("InvertTree() trees don't match")
			}
		})
	}
}
