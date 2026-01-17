package kth_smallest

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// KthSmallest returns the kth smallest value in a BST.
// Uses iterative in-order traversal for O(H + k) time complexity.
func KthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	current := root

	for len(stack) > 0 || current != nil {
		// Traverse to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop the next smallest node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return current.Val
		}

		// Move to right subtree
		current = current.Right
	}

	return -1 // Should never reach here if k is valid
}

// KthSmallestRecursive returns the kth smallest value using recursive in-order traversal.
func KthSmallestRecursive(root *TreeNode, k int) int {
	count := 0
	var result int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}

		// Traverse left subtree
		inorder(node.Left)

		// Visit current node
		count++
		if count == k {
			result = node.Val
			return
		}

		// Traverse right subtree
		inorder(node.Right)
	}

	inorder(root)
	return result
}
