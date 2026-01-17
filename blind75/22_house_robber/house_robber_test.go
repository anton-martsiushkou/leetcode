package house_robber

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
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
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
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "alternating high values",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
		{
			name: "increasing values",
			nums: []int{1, 2, 3, 4, 5},
			want: 9,
		},
		{
			name: "decreasing values",
			nums: []int{5, 4, 3, 2, 1},
			want: 9,
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
