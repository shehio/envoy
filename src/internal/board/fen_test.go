package board

import (
	"testing"
)

func TestSetFEN(t *testing.T) {
	board := NewBoard()
	tests := []struct {
		fen      string
		expected string
	}{
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
			"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
		},
		{
			"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
			"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
		},
	}

	for _, test := range tests {
		err := board.SetFEN(test.fen)
		if err != nil {
			t.Errorf("SetFEN(%s) failed: %v", test.fen, err)
			continue
		}
		if board.FEN() != test.expected {
			t.Errorf("SetFEN(%s) = %s; want %s", test.fen, board.FEN(), test.expected)
		}
	}
}

func TestInvalidFEN(t *testing.T) {
	board := NewBoard()
	tests := []struct {
		fen string
	}{
		{""},
		{"invalid"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"}, // Missing turn
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w"}, // Missing castling rights
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq"}, // Missing en passant
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"}, // Missing half move clock
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0"}, // Missing full move number
		{"rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"}, // Invalid rank length
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - -1 1"}, // Invalid half move clock
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 0"}, // Invalid full move number
	}

	for _, test := range tests {
		err := board.SetFEN(test.fen)
		if err == nil {
			t.Errorf("SetFEN(%s) should have failed", test.fen)
		}
	}
}

func TestFENRoundTrip(t *testing.T) {
	board := NewBoard()
	tests := []struct {
		fen string
	}{
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"},
		{"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"},
		{"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"},
	}

	for _, test := range tests {
		// Set FEN
		err := board.SetFEN(test.fen)
		if err != nil {
			t.Errorf("SetFEN(%s) failed: %v", test.fen, err)
			continue
		}

		// Get FEN
		fen := board.FEN()
		if fen != test.fen {
			t.Errorf("FEN() = %s; want %s", fen, test.fen)
			continue
		}

		// Set FEN again
		err = board.SetFEN(fen)
		if err != nil {
			t.Errorf("SetFEN(%s) failed on round trip: %v", fen, err)
			continue
		}

		// Get FEN again
		fen2 := board.FEN()
		if fen2 != test.fen {
			t.Errorf("FEN() = %s after round trip; want %s", fen2, test.fen)
		}
	}
}

func TestFEN(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Board)
		expected string
	}{
		{
			name:     "Starting position",
			setup:    func(b *Board) { b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1") },
			expected: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board := NewBoard()
			test.setup(board)
			result := board.FEN()
			if result != test.expected {
				t.Errorf("Expected FEN %s, got %s", test.expected, result)
			}
		})
	}
} 