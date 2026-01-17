package word_search_ii

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name: "example 1",
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"eat", "oath"},
		},
		{
			name: "example 2",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"abcb"},
			want:  []string{},
		},
		{
			name: "single cell board",
			board: [][]byte{
				{'a'},
			},
			words: []string{"a", "b"},
			want:  []string{"a"},
		},
		{
			name: "no matches",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"xyz", "pqr"},
			want:  []string{},
		},
		{
			name: "all words match",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"ab", "cd", "ac", "bd"},
			want:  []string{"ab", "cd", "ac", "bd"},
		},
		{
			name: "overlapping paths",
			board: [][]byte{
				{'a', 'a'},
				{'a', 'a'},
			},
			words: []string{"aa", "aaa", "aaaa"},
			want:  []string{"aa", "aaa", "aaaa"},
		},
		{
			name: "longer word",
			board: [][]byte{
				{'a', 'b', 'c'},
				{'d', 'e', 'f'},
				{'g', 'h', 'i'},
			},
			words: []string{"abef", "abeh", "abehi"},
			want:  []string{"abef", "abehi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the board to avoid modification issues
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}

			got := FindWords(boardCopy, tt.words)

			// Sort both slices for comparison
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindWordsEmptyBoard(t *testing.T) {
	board := [][]byte{}
	words := []string{"test"}
	got := FindWords(board, words)

	if len(got) != 0 {
		t.Errorf("FindWords() with empty board = %v, want []", got)
	}
}
