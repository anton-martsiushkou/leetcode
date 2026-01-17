package lca_bst

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	// Build the tree from example 1 and 2
	//        6
	//       / \
	//      2   8
	//     / \ / \
	//    0  4 7  9
	//      / \
	//     3   5
	root := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 8}
	node0 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node9 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}

	root.Left = node2
	root.Right = node8
	node2.Left = node0
	node2.Right = node4
	node8.Left = node7
	node8.Right = node9
	node4.Left = node3
	node4.Right = node5

	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			name: "example 1 - LCA is root",
			root: root,
			p:    node2,
			q:    node8,
			want: root,
		},
		{
			name: "example 2 - LCA is one of the nodes",
			root: root,
			p:    node2,
			q:    node4,
			want: node2,
		},
		{
			name: "both nodes in left subtree",
			root: root,
			p:    node0,
			q:    node4,
			want: node2,
		},
		{
			name: "both nodes in right subtree",
			root: root,
			p:    node7,
			q:    node9,
			want: node8,
		},
		{
			name: "deep nodes",
			root: root,
			p:    node3,
			q:    node5,
			want: node4,
		},
		{
			name: "example 3 - simple tree",
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 1},
			want: &TreeNode{Val: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowestCommonAncestor(tt.root, tt.p, tt.q)
			if got != nil && tt.want != nil && got.Val != tt.want.Val {
				t.Errorf("LowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}

func TestLowestCommonAncestorRecursive(t *testing.T) {
	// Build the tree
	root := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 8}
	node0 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}

	root.Left = node2
	root.Right = node8
	node2.Left = node0
	node2.Right = node4

	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			name: "LCA is root",
			root: root,
			p:    node2,
			q:    node8,
			want: root,
		},
		{
			name: "LCA is one of the nodes",
			root: root,
			p:    node2,
			q:    node4,
			want: node2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowestCommonAncestorRecursive(tt.root, tt.p, tt.q)
			if got != nil && tt.want != nil && got.Val != tt.want.Val {
				t.Errorf("LowestCommonAncestorRecursive() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
