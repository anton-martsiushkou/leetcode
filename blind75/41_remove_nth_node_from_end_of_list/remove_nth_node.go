package remove_nth_node

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd removes the nth node from the end using two pointers (one pass).
// Time complexity: O(n), Space complexity: O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy

	// Move fast n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node
	slow.Next = slow.Next.Next

	return dummy.Next
}

// RemoveNthFromEndTwoPass removes the nth node using two passes.
// Time complexity: O(n), Space complexity: O(1)
func RemoveNthFromEndTwoPass(head *ListNode, n int) *ListNode {
	// First pass: count the length
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	// Edge case: remove the first node
	if n == length {
		return head.Next
	}

	// Second pass: move to the node before the one to remove
	dummy := &ListNode{Next: head}
	curr = dummy
	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}

	// Remove the node
	curr.Next = curr.Next.Next

	return dummy.Next
}
