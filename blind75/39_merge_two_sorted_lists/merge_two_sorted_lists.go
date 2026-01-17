package merge_two_sorted_lists

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists iteratively.
// Time complexity: O(n + m), Space complexity: O(1)
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Attach the remaining nodes
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

// MergeTwoListsRecursive merges two sorted linked lists recursively.
// Time complexity: O(n + m), Space complexity: O(n + m) due to call stack
func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	}

	list2.Next = MergeTwoListsRecursive(list1, list2.Next)
	return list2
}
