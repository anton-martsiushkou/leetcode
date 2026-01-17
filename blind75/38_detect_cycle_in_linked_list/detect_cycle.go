package detect_cycle

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle determines if a linked list has a cycle using Floyd's algorithm.
// Time complexity: O(n), Space complexity: O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// HasCycleHashSet determines if a linked list has a cycle using a hash set.
// Time complexity: O(n), Space complexity: O(n)
func HasCycleHashSet(head *ListNode) bool {
	seen := make(map[*ListNode]bool)

	curr := head
	for curr != nil {
		if seen[curr] {
			return true
		}
		seen[curr] = true
		curr = curr.Next
	}

	return false
}
