package validate_bst

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "example 1 - valid BST",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "example 2 - invalid BST",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "invalid - left child greater than root",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "invalid - deep violation",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 15}, // 15 > 10, violates BST
				},
				Right: &TreeNode{Val: 20},
			},
			want: false,
		},
		{
			name: "valid - all same side",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			name: "invalid - duplicate values",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBST(tt.root)
			if got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
