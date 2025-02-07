package main

import (
	"testing"
)

func TestNewChessCoordinator(t *testing.T) {
	tests := []struct {
		name        string
		whitePlayerURL  string
		blackPlayerURL  string
		expectError bool
	}{
		{
			name:        "Valid URLs",
			whitePlayerURL:  "http://localhost:8081",
			blackPlayerURL:  "http://localhost:8082",
			expectError: false,
		},
		{
			name:        "Empty URLs",
			whitePlayerURL:  "",
			blackPlayerURL:  "",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			coordinator := NewChessCoordinator(test.whitePlayerURL, test.blackPlayerURL)

			if coordinator == nil {
				t.Error("Coordinator is nil")
			}

			if coordinator.whitePlayerURL != test.whitePlayerURL {
				t.Errorf("Expected whitePlayerURL %s, got %s", test.whitePlayerURL, coordinator.whitePlayerURL)
			}

			if coordinator.blackPlayerURL != test.blackPlayerURL {
				t.Errorf("Expected blackPlayerURL %s, got %s", test.blackPlayerURL, coordinator.blackPlayerURL)
			}

			if coordinator.board == nil {
				t.Error("Board is nil")
			}
		})
	}
}

func TestGetMoveFromPlayer(t *testing.T) {
	coordinator := NewChessCoordinator("http://localhost:8081", "http://localhost:8082")

	tests := []struct {
		name        string
		fen         string
		expectError bool
	}{
		{
			name:        "Valid request",
			fen:         "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectError: true, // Will fail because no server is running
		},
		{
			name:        "Invalid FEN",
			fen:         "invalid",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			move, err := coordinator.getMoveFromPlayer(test.fen)
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

func TestMakeMove(t *testing.T) {
	coordinator := NewChessCoordinator("http://localhost:8081", "http://localhost:8082")

	tests := []struct {
		name        string
		move        string
		expectError bool
	}{
		{
			name:        "Valid pawn move",
			move:        "e2e4",
			expectError: false,
		},
		{
			name:        "Invalid move",
			move:        "invalid",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := coordinator.makeMove(test.move)
			if test.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Verify the move was made correctly
			if coordinator.board.IsWhiteToMove() {
				t.Error("Expected turn to switch to black")
			}
		})
	}
} 