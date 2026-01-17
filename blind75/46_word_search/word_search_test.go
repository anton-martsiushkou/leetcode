package word_search

import (
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "example 1 - word exists",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "example 2 - word exists",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			name: "example 3 - word does not exist",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		{
			name: "single cell match",
			board: [][]byte{
				{'A'},
			},
			word: "A",
			want: true,
		},
		{
			name: "single cell no match",
			board: [][]byte{
				{'A'},
			},
			word: "B",
			want: false,
		},
		{
			name: "word longer than board",
			board: [][]byte{
				{'A', 'B'},
			},
			word: "ABCD",
			want: false,
		},
		{
			name: "requires backtracking",
			board: [][]byte{
				{'C', 'A', 'A'},
				{'A', 'A', 'A'},
				{'B', 'C', 'D'},
			},
			word: "AAB",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the board to avoid modifying the test case
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}

			got := Exist(boardCopy, tt.word)
			if got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
