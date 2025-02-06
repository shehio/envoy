package board

import (
	"testing"
)

func TestIsGameOver(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		expected bool
	}{
		{
			name:     "Game not over - starting position",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expected: false,
		},
		{
			name:     "Game not over - middle game",
			fen:      "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2",
			expected: false,
		},
		{
			name:     "Game over - white king captured",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQ1BNR w KQkq - 0 1",
			expected: true,
		},
		{
			name:     "Game over - black king captured",
			fen:      "rnbq1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expected: true,
		},
		{
			name:     "Game over - both kings captured",
			fen:      "rnbq1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQ1BNR w KQkq - 0 1",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board := NewBoard()
			err := board.SetFEN(test.fen)
			if err != nil {
				t.Fatalf("Failed to set FEN: %v", err)
			}

			result := board.IsGameOver()
			if result != test.expected {
				t.Errorf("IsGameOver() = %v, expected %v", result, test.expected)
			}
		})
	}
}

func TestResult(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		expected string
	}{
		{
			name:     "Game not over - starting position",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expected: "*",
		},
		{
			name:     "Game not over - middle game",
			fen:      "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2",
			expected: "*",
		},
		{
			name:     "Game over - white king captured",
			fen:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQ1BNR w KQkq - 0 1",
			expected: "0-1",
		},
		{
			name:     "Game over - black king captured",
			fen:      "rnbq1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expected: "1-0",
		},
		{
			name:     "Game over - both kings captured",
			fen:      "rnbq1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQ1BNR w KQkq - 0 1",
			expected: "1/2-1/2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board := NewBoard()
			err := board.SetFEN(test.fen)
			if err != nil {
				t.Fatalf("Failed to set FEN: %v", err)
			}

			result := board.Result()
			if result != test.expected {
				t.Errorf("Result() = %s, expected %s", result, test.expected)
			}
		})
	}
} 