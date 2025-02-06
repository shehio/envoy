package stockfish

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}
	defer handler.Close()

	if handler == nil {
		t.Error("Handler is nil")
	}
}

func TestGetMove(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}
	defer handler.Close()

	tests := []struct {
		name        string
		fen         string
		expectError bool
	}{
		{
			name:        "Starting position",
			fen:         "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectError: false,
		},
		{
			name:        "Invalid FEN",
			fen:         "invalid",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			move, err := handler.GetMove(test.fen)
			if test.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if move == "" {
				t.Error("Expected a move but got empty string")
			}

			// Verify move format (e.g., "e2e4")
			if len(move) != 4 {
				t.Errorf("Invalid move format: %s", move)
			}
		})
	}
} 