package counting_bits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "example 1",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "example 2",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "n is 0",
			n:    0,
			want: []int{0},
		},
		{
			name: "n is 1",
			n:    1,
			want: []int{0, 1},
		},
		{
			name: "n is 8",
			n:    8,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			name: "n is 15",
			n:    15,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountBits(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
