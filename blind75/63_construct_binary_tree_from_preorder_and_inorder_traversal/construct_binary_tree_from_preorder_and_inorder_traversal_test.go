package construct_binary_tree_from_preorder_and_inorder_traversal

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

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "example 1 - balanced tree",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name:     "example 2 - single node",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode{Val: -1},
		},
		{
			name:     "left skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
		{
			name:     "right skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
		},
		{
			name:     "complete binary tree",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			want: &TreeNode{
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
		},
		{
			name:     "unbalanced tree",
			preorder: []int{1, 2, 4, 3},
			inorder:  []int{4, 2, 1, 3},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name:     "two nodes",
			preorder: []int{1, 2},
			inorder:  []int{2, 1},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTree(tt.preorder, tt.inorder)
			if !isSameTree(got, tt.want) {
				t.Errorf("BuildTree() trees don't match")
			}
		})
	}
}
