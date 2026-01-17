package number_of_islands

import (
	"testing"
)

func copyGrid(grid [][]byte) [][]byte {
	copied := make([][]byte, len(grid))
	for i := range grid {
		copied[i] = make([]byte, len(grid[i]))
		copy(copied[i], grid[i])
	}
	return copied
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example 1 - single island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "example 2 - three islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "no islands",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all islands",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "single cell island",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "diagonal islands",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies the grid
			gridCopy := copyGrid(tt.grid)
			got := NumIslands(gridCopy)
			if got != tt.want {
				t.Errorf("NumIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumIslandsBFS(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example 1 - single island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "example 2 - three islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "all islands",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "diagonal islands",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := copyGrid(tt.grid)
			got := NumIslandsBFS(gridCopy)
			if got != tt.want {
				t.Errorf("NumIslandsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
