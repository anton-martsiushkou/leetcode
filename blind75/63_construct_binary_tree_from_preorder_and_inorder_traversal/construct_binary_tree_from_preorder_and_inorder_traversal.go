package construct_binary_tree_from_preorder_and_inorder_traversal

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from preorder and inorder traversals.
// Uses recursive DFS with hashmap for O(n) time complexity.
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// Create hashmap for O(1) lookup of inorder indices
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	preorderIndex := 0

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		// Base case: no elements to construct the tree
		if left > right {
			return nil
		}

		// Select the preorder index element as root and increment
		rootVal := preorder[preorderIndex]
		preorderIndex++

		root := &TreeNode{Val: rootVal}

		// Build left and right subtree
		// excluding inorderMap[rootVal] element because it's the root
		root.Left = build(left, inorderMap[rootVal]-1)
		root.Right = build(inorderMap[rootVal]+1, right)

		return root
	}

	return build(0, len(inorder)-1)
}
