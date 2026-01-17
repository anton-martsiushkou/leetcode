package kth_smallest

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4},
			},
			k:    1,
			want: 1,
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:    3,
			want: 3,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			k:    1,
			want: 1,
		},
		{
			name: "k equals tree size",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			k:    3,
			want: 3,
		},
		{
			name: "left-skewed tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
			k:    2,
			want: 2,
		},
		{
			name: "right-skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("KthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKthSmallestRecursive(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4},
			},
			k:    1,
			want: 1,
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthSmallestRecursive(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("KthSmallestRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
