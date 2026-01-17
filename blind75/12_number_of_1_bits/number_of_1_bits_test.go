package number_of_1_bits

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    11,
			want: 3,
		},
		{
			name: "example 2",
			n:    128,
			want: 1,
		},
		{
			name: "example 3",
			n:    2147483645,
			want: 30,
		},
		{
			name: "single bit set",
			n:    1,
			want: 1,
		},
		{
			name: "all bits set in 8-bit",
			n:    255,
			want: 8,
		},
		{
			name: "power of two",
			n:    1024,
			want: 1,
		},
		{
			name: "alternating bits",
			n:    21845, // 0101010101010101
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HammingWeight(tt.n)
			if got != tt.want {
				t.Errorf("HammingWeight(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
