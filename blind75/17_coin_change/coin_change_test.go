package coin_change

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "example 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "example 2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "example 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "greedy fails",
			coins:  []int{1, 3, 4},
			amount: 6,
			want:   2,
		},
		{
			name:   "large amount",
			coins:  []int{1, 5, 10, 25},
			amount: 100,
			want:   4,
		},
		{
			name:   "impossible with large coins",
			coins:  []int{5, 10},
			amount: 3,
			want:   -1,
		},
		{
			name:   "single coin exact match",
			coins:  []int{7},
			amount: 7,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
