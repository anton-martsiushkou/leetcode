package reorder_list

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReorderList reorders the list in-place using find middle + reverse + merge.
// Time complexity: O(n), Space complexity: O(1)
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find the middle of the list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Split and reverse the second half
	second := slow.Next
	slow.Next = nil // Split the list
	second = reverseList(second)

	// Step 3: Merge the two halves
	first := head
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}

// reverseList reverses a linked list
func reverseList(head *ListNode) *ListNode {
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

// ReorderListStack reorders the list using a stack.
// Time complexity: O(n), Space complexity: O(n)
func ReorderListStack(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Push all nodes to stack
	var stack []*ListNode
	curr := head
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Next
	}

	// Reorder by taking from front and back alternately
	curr = head
	n := len(stack)
	for i := 0; i < n/2; i++ {
		next := curr.Next
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr.Next = back
		back.Next = next

		curr = next
	}

	// Set the end of the list
	if curr != nil {
		curr.Next = nil
	}
}
