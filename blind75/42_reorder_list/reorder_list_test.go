package reorder_list

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

func TestReorderList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "example 1 - even length",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 4, 2, 3},
		},
		{
			name:  "example 2 - odd length",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 5, 2, 4, 3},
		},
		{
			name:  "single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "two nodes",
			input: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "three nodes",
			input: []int{1, 2, 3},
			want:  []int{1, 3, 2},
		},
		{
			name:  "six nodes",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{1, 6, 2, 5, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			ReorderList(head)
			result := listToSlice(head)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReorderList() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestReorderListStack(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "example 1 - even length",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 4, 2, 3},
		},
		{
			name:  "example 2 - odd length",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 5, 2, 4, 3},
		},
		{
			name:  "single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "two nodes",
			input: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "three nodes",
			input: []int{1, 2, 3},
			want:  []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			ReorderListStack(head)
			result := listToSlice(head)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReorderListStack() = %v, want %v", result, tt.want)
			}
		})
	}
}
