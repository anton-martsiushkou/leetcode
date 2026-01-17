package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec handles serialization and deserialization of binary trees.
type Codec struct{}

// Constructor creates a new Codec instance.
func Constructor() Codec {
	return Codec{}
}

// Serialize encodes a tree to a single string using preorder traversal.
func (c *Codec) Serialize(root *TreeNode) string {
	values := []string{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			values = append(values, "null")
			return
		}

		values = append(values, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(values, ",")
}

// Deserialize decodes your encoded data to tree.
func (c *Codec) Deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(values) || values[index] == "null" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(values[index])
		index++

		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}
