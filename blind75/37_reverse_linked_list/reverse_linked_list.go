package reverse_linked_list

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a singly linked list using an iterative approach.
// Time complexity: O(n), Space complexity: O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// ReverseListRecursive reverses a singly linked list using a recursive approach.
// Time complexity: O(n), Space complexity: O(n) due to call stack
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
