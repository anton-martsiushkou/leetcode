package reverse_bits

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{
			name: "example 1",
			n:    0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "example 2",
			n:    0b11111111111111111111111111111101,
			want: 0b10111111111111111111111111111111,
		},
		{
			name: "all zeros",
			n:    0b00000000000000000000000000000000,
			want: 0b00000000000000000000000000000000,
		},
		{
			name: "all ones",
			n:    0b11111111111111111111111111111111,
			want: 0b11111111111111111111111111111111,
		},
		{
			name: "single bit set at start",
			n:    0b10000000000000000000000000000000,
			want: 0b00000000000000000000000000000001,
		},
		{
			name: "single bit set at end",
			n:    0b00000000000000000000000000000001,
			want: 0b10000000000000000000000000000000,
		},
		{
			name: "alternating bits",
			n:    0b10101010101010101010101010101010,
			want: 0b01010101010101010101010101010101,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseBits(tt.n)
			if got != tt.want {
				t.Errorf("ReverseBits(%b) = %b, want %b", tt.n, got, tt.want)
			}
		})
	}
}
