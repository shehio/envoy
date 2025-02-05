package board

import (
	"fmt"
	"testing"
)

func TestParseSquare(t *testing.T) {
	tests := []struct {
		square   string
		expected int
	}{
		{"a1", 0},
		{"b1", 1},
		{"c1", 2},
		{"d1", 3},
		{"e1", 4},
		{"f1", 5},
		{"g1", 6},
		{"h1", 7},
		{"a2", 8},
		{"b2", 9},
		{"c2", 10},
		{"d2", 11},
		{"e2", 12},
		{"f2", 13},
		{"g2", 14},
		{"h2", 15},
		{"a8", 56},
		{"b8", 57},
		{"c8", 58},
		{"d8", 59},
		{"e8", 60},
		{"f8", 61},
		{"g8", 62},
		{"h8", 63},
		{"invalid", -1},
		{"", -1},
		{"a9", -1},
		{"i1", -1},
		{"1a", -1},
	}

	for _, test := range tests {
		result := squareToIndex(test.square)
		if result != test.expected {
			t.Errorf("squareToIndex(%s) = %d; want %d", test.square, result, test.expected)
		}
	}
}

func TestSquareToIndex(t *testing.T) {
	tests := []struct {
		square   string
		expected int
	}{
		{"a1", 0},
		{"b1", 1},
		{"c1", 2},
		{"d1", 3},
		{"e1", 4},
		{"f1", 5},
		{"g1", 6},
		{"h1", 7},
		{"a2", 8},
		{"b2", 9},
		{"c2", 10},
		{"d2", 11},
		{"e2", 12},
		{"f2", 13},
		{"g2", 14},
		{"h2", 15},
		{"a8", 56},
		{"b8", 57},
		{"c8", 58},
		{"d8", 59},
		{"e8", 60},
		{"f8", 61},
		{"g8", 62},
		{"h8", 63},
		{"invalid", -1},
		{"", -1},
		{"a9", -1},
		{"i1", -1},
		{"1a", -1},
	}

	for _, test := range tests {
		result := squareToIndex(test.square)
		if result != test.expected {
			t.Errorf("squareToIndex(%s) = %d; want %d", test.square, result, test.expected)
		}
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()
	tests := []struct {
		name     string
		move     Move
		expected error
	}{
		{
			name:     "Valid pawn move",
			move:     Move{From: "e2", To: "e4"},
			expected: nil,
		},
		{
			name:     "Invalid square",
			move:     Move{From: "invalid", To: "e4"},
			expected: fmt.Errorf("invalid square"),
		},
		{
			name:     "No piece at source",
			move:     Move{From: "e3", To: "e4"},
			expected: fmt.Errorf("no piece at source square"),
		},
		{
			name:     "Wrong player's turn",
			move:     Move{From: "b1", To: "c3"},
			expected: fmt.Errorf("not your turn"),
		},
		{
			name:     "Valid promotion",
			move:     Move{From: "e6", To: "e8", Promotion: BlackQueen},
			expected: fmt.Errorf("invalid promotion"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "Wrong player's turn" {
				// Make a move to switch to black's turn
				board.MakeMove(Move{From: "e2", To: "e4"})
			} else if test.name == "Valid promotion" {
				// Set up a position where black can promote
				board.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
				// Move the black pawn to e6
				board.MakeMove(Move{From: "e7", To: "e6"})
				// Move a white piece to allow black to move again
				board.MakeMove(Move{From: "e2", To: "e4"})
			}
			err := board.MakeMove(test.move)
			if err == nil && test.expected != nil {
				t.Errorf("MakeMove(%v) should have failed", test.move)
			} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
				t.Errorf("MakeMove(%v) = %v; want %v", test.move, err, test.expected)
			}
		})
	}
}

func TestMoveString(t *testing.T) {
	tests := []struct {
		move     Move
		expected string
	}{
		{Move{From: "e2", To: "e4"}, "e2e4"},
		{Move{From: "b1", To: "c3"}, "b1c3"},
		{Move{From: "e1", To: "g1"}, "e1g1"},
		{Move{From: "e7", To: "e8", Promotion: BlackQueen}, "e7e8q"},
	}

	for _, test := range tests {
		result := test.move.String()
		if result != test.expected {
			t.Errorf("Move.String() = %s; want %s", result, test.expected)
		}
	}
} 