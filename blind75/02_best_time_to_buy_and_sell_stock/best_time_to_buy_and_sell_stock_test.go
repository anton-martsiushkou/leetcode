package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "example 2 - descending prices",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "ascending prices",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "single price",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "two prices - profit",
			prices: []int{1, 5},
			want:   4,
		},
		{
			name:   "two prices - no profit",
			prices: []int{5, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
