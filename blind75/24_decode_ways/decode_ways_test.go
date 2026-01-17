package decode_ways

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "12",
			want: 2,
		},
		{
			name: "example 2",
			s:    "226",
			want: 3,
		},
		{
			name: "example 3",
			s:    "06",
			want: 0,
		},
		{
			name: "single digit",
			s:    "1",
			want: 1,
		},
		{
			name: "leading zero",
			s:    "0",
			want: 0,
		},
		{
			name: "valid two digit",
			s:    "10",
			want: 1,
		},
		{
			name: "invalid zero in middle",
			s:    "100",
			want: 0,
		},
		{
			name: "multiple valid decodings",
			s:    "111111",
			want: 13,
		},
		{
			name: "boundary case 26",
			s:    "2611055971756562",
			want: 4,
		},
		{
			name: "single character valid",
			s:    "5",
			want: 1,
		},
		{
			name: "two digits out of range",
			s:    "27",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumDecodings(tt.s)
			if got != tt.want {
				t.Errorf("NumDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
