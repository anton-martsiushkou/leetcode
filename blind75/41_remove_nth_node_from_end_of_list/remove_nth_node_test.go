package remove_nth_node

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

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "example 1",
			input: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "example 2 - single node",
			input: []int{1},
			n:     1,
			want:  []int{},
		},
		{
			name:  "example 3 - remove last",
			input: []int{1, 2},
			n:     1,
			want:  []int{1},
		},
		{
			name:  "remove first",
			input: []int{1, 2},
			n:     2,
			want:  []int{2},
		},
		{
			name:  "remove middle",
			input: []int{1, 2, 3},
			n:     2,
			want:  []int{1, 3},
		},
		{
			name:  "remove from longer list",
			input: []int{1, 2, 3, 4, 5, 6},
			n:     3,
			want:  []int{1, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			got := RemoveNthFromEnd(head, tt.n)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestRemoveNthFromEndTwoPass(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "example 1",
			input: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "example 2 - single node",
			input: []int{1},
			n:     1,
			want:  []int{},
		},
		{
			name:  "example 3 - remove last",
			input: []int{1, 2},
			n:     1,
			want:  []int{1},
		},
		{
			name:  "remove first",
			input: []int{1, 2},
			n:     2,
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			got := RemoveNthFromEndTwoPass(head, tt.n)
			result := listToSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("RemoveNthFromEndTwoPass() = %v, want %v", result, tt.want)
			}
		})
	}
}
