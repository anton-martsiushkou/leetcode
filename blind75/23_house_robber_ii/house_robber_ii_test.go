package house_robber_ii

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example 3",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two houses",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "all same values",
			nums: []int{5, 5, 5, 5},
			want: 10,
		},
		{
			name: "increasing values",
			nums: []int{1, 2, 3, 4, 5},
			want: 8,
		},
		{
			name: "large first and last",
			nums: []int{100, 1, 1, 1, 100},
			want: 100,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob(tt.nums)
			if got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
