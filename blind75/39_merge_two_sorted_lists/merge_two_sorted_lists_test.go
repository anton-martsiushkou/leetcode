package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "example 1",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "example 2 - both empty",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "example 3 - first empty",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "second empty",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "no overlap",
			list1: []int{1, 2, 3},
			list2: []int{4, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "interleaved",
			list1: []int{1, 3, 5},
			list2: []int{2, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.list1)
			l2 := createList(tt.list2)
			got := MergeTwoLists(l1, l2)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestMergeTwoListsRecursive(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "example 1",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "example 2 - both empty",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "example 3 - first empty",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "second empty",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.list1)
			l2 := createList(tt.list2)
			got := MergeTwoListsRecursive(l1, l2)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("MergeTwoListsRecursive() = %v, want %v", result, tt.want)
			}
		})
	}
}
