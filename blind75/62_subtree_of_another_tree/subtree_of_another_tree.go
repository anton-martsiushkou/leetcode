package subtree_of_another_tree

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSubtree checks if subRoot is a subtree of root.
// Uses recursive DFS for O(m * n) time complexity.
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	// Check if current tree matches subRoot
	if isSameTree(root, subRoot) {
		return true
	}

	// Recursively check left and right subtrees
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

// isSameTree checks if two trees are identical.
func isSameTree(p *TreeNode, q *TreeNode) bool {
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
