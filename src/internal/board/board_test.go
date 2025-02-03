package board

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Error("NewBoard() returned nil")
	}
	if len(board.squares) != 64 {
		t.Errorf("Board size = %d; want 64", len(board.squares))
	}
}


func TestIsWhiteTurn(t *testing.T) {
	board := NewBoard()
	if !board.IsWhiteTurn() {
		t.Error("New board should have white's turn")
	}
}

func TestGetEnPassantSquare(t *testing.T) {
	board := NewBoard()
	if board.GetEnPassantSquare() != "-" {
		t.Errorf("GetEnPassantSquare() = %s; want -", board.GetEnPassantSquare())
	}
}

func TestSetEnPassantSquare(t *testing.T) {
	board := NewBoard()
	tests := []struct {
		square   string
		expected string
	}{
		{"e3", "e3"},
		{"d6", "d6"},
		{"-", "-"},
	}

	for _, test := range tests {
		board.SetEnPassantSquare(test.square)
		if board.GetEnPassantSquare() != test.expected {
			t.Errorf("SetEnPassantSquare(%s) = %s; want %s", test.square, board.GetEnPassantSquare(), test.expected)
		}
	}
}
