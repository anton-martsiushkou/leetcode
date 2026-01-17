package subtree_of_another_tree

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    *TreeNode
		subRoot *TreeNode
		want    bool
	}{
		{
			name: "example 1 - subtree exists",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			want: true,
		},
		{
			name: "example 2 - subtree has extra node",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 0},
					},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "example 3 - simple match",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
			},
			subRoot: &TreeNode{Val: 1},
			want:    true,
		},
		{
			name: "identical trees",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			subRoot: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "subRoot is leaf of root",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			subRoot: &TreeNode{Val: 4},
			want:    true,
		},
		{
			name: "no match - different values",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			subRoot: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
			want: false,
		},
		{
			name: "no match - different structure",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			subRoot: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "empty root",
			root: nil,
			subRoot: &TreeNode{Val: 1},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSubtree(tt.root, tt.subRoot)
			if got != tt.want {
				t.Errorf("IsSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
