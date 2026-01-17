package lca_bst

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor finds the LCA of two nodes in a BST.
// Uses iterative approach with O(H) time and O(1) space complexity.
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		// Both nodes are in left subtree
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
		} else if p.Val > current.Val && q.Val > current.Val {
			// Both nodes are in right subtree
			current = current.Right
		} else {
			// Split point found - one node is on left, other on right
			// (or current node equals one of them)
			return current
		}
	}

	return nil
}

// LowestCommonAncestorRecursive finds the LCA using recursive approach.
func LowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Both nodes are in left subtree
	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestorRecursive(root.Left, p, q)
	}

	// Both nodes are in right subtree
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestorRecursive(root.Right, p, q)
	}

	// Split point found
	return root
}
