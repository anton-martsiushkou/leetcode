package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
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
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "example 2 - single node",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "example 3 - empty tree",
			root: nil,
			want: [][]int{},
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
			want: [][]int{{1}, {2}, {3}},
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
			want: [][]int{{1}, {2}, {3}},
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
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			want: [][]int{{1}, {2, 3}, {4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrder(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
