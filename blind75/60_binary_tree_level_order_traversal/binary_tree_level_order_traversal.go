package binary_tree_level_order_traversal

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder returns the level order traversal of a binary tree.
// Uses BFS with a queue for O(n) time complexity.
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}
