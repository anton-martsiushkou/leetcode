package same_tree

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSameTree checks if two binary trees are identical.
// Uses recursive DFS for O(n) time complexity.
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// Both nodes are nil - same structure
	if p == nil && q == nil {
		return true
	}

	// One node is nil, the other isn't - different structure
	if p == nil || q == nil {
		return false
	}

	// Values differ - not the same
	if p.Val != q.Val {
		return false
	}

	// Check if left and right subtrees are the same
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
