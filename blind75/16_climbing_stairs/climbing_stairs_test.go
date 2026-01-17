package climbing_stairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    2,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			want: 3,
		},
		{
			name: "example 3",
			n:    4,
			want: 5,
		},
		{
			name: "base case n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=5",
			n:    5,
			want: 8,
		},
		{
			name: "n=10",
			n:    10,
			want: 89,
		},
		{
			name: "large n=20",
			n:    20,
			want: 10946,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.n)
			if got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
