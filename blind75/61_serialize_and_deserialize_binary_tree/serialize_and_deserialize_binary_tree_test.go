package serialize_and_deserialize_binary_tree

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

func TestCodec(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "example 1 - complete tree",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			name: "example 2 - empty tree",
			root: nil,
		},
		{
			name: "example 3 - single node",
			root: &TreeNode{Val: 1},
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
		},
		{
			name: "complex tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 7},
					},
				},
			},
		},
		{
			name: "tree with negative values",
			root: &TreeNode{
				Val:   -1,
				Left:  &TreeNode{Val: -2},
				Right: &TreeNode{Val: -3},
			},
		},
		{
			name: "large values",
			root: &TreeNode{
				Val:   1000,
				Left:  &TreeNode{Val: -1000},
				Right: &TreeNode{Val: 500},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()
			serialized := codec.Serialize(tt.root)
			deserialized := codec.Deserialize(serialized)

			if !isSameTree(tt.root, deserialized) {
				t.Errorf("Serialize/Deserialize failed: trees don't match")
			}
		})
	}
}

func TestSerializeFormat(t *testing.T) {
	codec := Constructor()

	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: "1,null,null",
		},
		{
			name: "empty tree",
			root: nil,
			want: "null",
		},
		{
			name: "simple tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: "1,2,null,null,3,null,null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := codec.Serialize(tt.root)
			if got != tt.want {
				t.Errorf("Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
