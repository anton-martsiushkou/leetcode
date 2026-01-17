package merge_k_sorted_lists

import "container/heap"

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap implements heap.Interface for ListNode pointers
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKLists merges k sorted linked lists using a min heap.
// Time complexity: O(N log k), Space complexity: O(k)
// where N is the total number of nodes and k is the number of lists
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// Initialize min heap with first node from each list
	h := &NodeHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	// Build merged list
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = tail.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// MergeKListsDivideConquer merges k sorted lists using divide and conquer.
// Time complexity: O(N log k), Space complexity: O(log k) for recursion
func MergeKListsDivideConquer(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	l1 := mergeHelper(lists, left, mid)
	l2 := mergeHelper(lists, mid+1, right)

	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
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

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}
