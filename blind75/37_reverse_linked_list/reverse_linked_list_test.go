package reverse_linked_list

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

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "example 1",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "example 2",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "example 3 - empty list",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "three nodes",
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			got := ReverseList(head)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "example 1",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "example 2",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "example 3 - empty list",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "three nodes",
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			got := ReverseListRecursive(head)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReverseListRecursive() = %v, want %v", result, tt.want)
			}
		})
	}
}
