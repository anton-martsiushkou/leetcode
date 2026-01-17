package invert_binary_tree

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InvertTree inverts a binary tree (mirrors it).
// Uses recursive DFS for O(n) time complexity.
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert subtrees
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
