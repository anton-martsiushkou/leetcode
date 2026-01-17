package same_tree

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "example 1 - identical trees",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "example 2 - different structure",
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "example 3 - different values",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "both empty",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "one empty",
			p:    &TreeNode{Val: 1},
			q:    nil,
			want: false,
		},
		{
			name: "single node same",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 5},
			want: true,
		},
		{
			name: "single node different",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 10},
			want: false,
		},
		{
			name: "deep identical trees",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 6},
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 6},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
