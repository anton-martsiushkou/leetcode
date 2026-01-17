package detect_cycle

import (
	"testing"
)

// Helper function to create a linked list with a cycle
func createListWithCycle(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	var cycleNode *ListNode

	if pos == 0 {
		cycleNode = head
	}

	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
		if i == pos {
			cycleNode = curr
		}
	}

	// Create the cycle
	if pos >= 0 {
		curr.Next = cycleNode
	}

	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{
			name: "example 1 - cycle at position 1",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			name: "example 2 - cycle at position 0",
			vals: []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			name: "example 3 - no cycle",
			vals: []int{1},
			pos:  -1,
			want: false,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
			want: false,
		},
		{
			name: "no cycle - multiple nodes",
			vals: []int{1, 2, 3, 4, 5},
			pos:  -1,
			want: false,
		},
		{
			name: "cycle at end",
			vals: []int{1, 2, 3},
			pos:  2,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createListWithCycle(tt.vals, tt.pos)
			got := HasCycle(head)
			if got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycleHashSet(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{
			name: "example 1 - cycle at position 1",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			name: "example 2 - cycle at position 0",
			vals: []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			name: "example 3 - no cycle",
			vals: []int{1},
			pos:  -1,
			want: false,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
			want: false,
		},
		{
			name: "no cycle - multiple nodes",
			vals: []int{1, 2, 3, 4, 5},
			pos:  -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createListWithCycle(tt.vals, tt.pos)
			got := HasCycleHashSet(head)
			if got != tt.want {
				t.Errorf("HasCycleHashSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
