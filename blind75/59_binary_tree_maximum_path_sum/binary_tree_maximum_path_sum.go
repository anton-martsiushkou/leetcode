package binary_tree_maximum_path_sum

import "math"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxPathSum returns the maximum path sum in a binary tree.
// Uses recursive DFS with global maximum tracking for O(n) time complexity.
func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Get maximum gain from left and right subtrees (ignore negative gains)
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// Calculate path sum through current node
		pathSum := node.Val + leftGain + rightGain

		// Update global maximum
		maxSum = max(maxSum, pathSum)

		// Return maximum single path (can only use one child)
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
