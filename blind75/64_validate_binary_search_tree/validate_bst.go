package validate_bst

import "math"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsValidBST determines if a binary tree is a valid BST.
// Uses recursive range validation for O(n) time complexity.
func IsValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

// isValidBSTHelper validates that the node's value is within the valid range
// and recursively validates left and right subtrees with updated ranges.
func isValidBSTHelper(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	// Check if current node's value is within valid range
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}

	// Validate left subtree (all values must be < node.Val)
	// and right subtree (all values must be > node.Val)
	return isValidBSTHelper(node.Left, min, int64(node.Val)) &&
		isValidBSTHelper(node.Right, int64(node.Val), max)
}
