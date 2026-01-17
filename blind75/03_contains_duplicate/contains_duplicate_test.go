package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1 - has duplicate",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "example 2 - no duplicate",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "example 3 - multiple duplicates",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "two same elements",
			nums: []int{1, 1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
