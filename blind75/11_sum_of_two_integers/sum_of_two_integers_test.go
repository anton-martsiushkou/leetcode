package sum_of_two_integers

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "example 1",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "example 2",
			a:    2,
			b:    3,
			want: 5,
		},
		{
			name: "example 3 - negative and positive",
			a:    -1,
			b:    1,
			want: 0,
		},
		{
			name: "both positive",
			a:    5,
			b:    3,
			want: 8,
		},
		{
			name: "both negative",
			a:    -5,
			b:    -3,
			want: -8,
		},
		{
			name: "zero and positive",
			a:    0,
			b:    5,
			want: 5,
		},
		{
			name: "zero and negative",
			a:    0,
			b:    -5,
			want: -5,
		},
		{
			name: "both zero",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("GetSum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
