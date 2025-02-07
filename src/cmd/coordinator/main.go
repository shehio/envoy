package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shehio/envoy/src/internal/board"
	"github.com/shehio/envoy/pkg/types"
)

type ChessCoordinator struct {
	whitePlayerURL string
	blackPlayerURL string
	board         *board.Board
}

func NewChessCoordinator(whitePlayerURL, blackPlayerURL string) *ChessCoordinator {
	return &ChessCoordinator{
		whitePlayerURL: whitePlayerURL,
		blackPlayerURL: blackPlayerURL,
		board:         board.NewBoard(),
	}
}

func (c *ChessCoordinator) getMoveFromPlayer(fen string) (string, error) {
	url := c.whitePlayerURL
	if !c.board.IsWhiteToMove() {
		url = c.blackPlayerURL
	}

	req := types.MoveRequest{FEN: fen}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("player returned status %d", resp.StatusCode)
	}

	var moveResp types.MoveResponse
	if err := json.NewDecoder(resp.Body).Decode(&moveResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return moveResp.Move, nil
}

func (c *ChessCoordinator) makeMove(moveStr string) error {
	move, err := board.ParseMove(moveStr)
	if err != nil {
		return fmt.Errorf("invalid move format: %v", err)
	}
	return c.board.MakeMove(move)
}

// ASCII representation of the board
func fenToASCII(fen string) string {
	parts := strings.Fields(fen)
	if len(parts) < 1 {
		return "Invalid FEN string"
	}

	// Split the position part into ranks
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return "Invalid FEN string"
	}

	var board strings.Builder
	board.WriteString("  a b c d e f g h\n")
	board.WriteString("  ---------------\n")

	// Process each rank from 8 to 1
	for i := 7; i >= 0; i-- {
		board.WriteString(fmt.Sprintf("%d|", i+1))
		rank := ranks[i]
		for _, char := range rank {
			if char >= '1' && char <= '8' {
				spaces := int(char - '0')
				board.WriteString(strings.Repeat("  ", spaces))
			} else {
				// Convert piece characters to more readable symbols
				piece := string(char)
				switch piece {
				case "p": piece = "♟"
				case "P": piece = "♙"
				case "r": piece = "♜"
				case "R": piece = "♖"
				case "n": piece = "♞"
				case "N": piece = "♘"
				case "b": piece = "♝"
				case "B": piece = "♗"
				case "q": piece = "♛"
				case "Q": piece = "♕"
				case "k": piece = "♚"
				case "K": piece = "♔"
				}
				board.WriteString(piece + " ")
			}
		}
		board.WriteString(fmt.Sprintf("|%d\n", i+1))
	}
	board.WriteString("  ---------------\n")
	board.WriteString("  a b c d e f g h\n")

	// Add game state information
	if len(parts) > 1 {
		board.WriteString("\nGame State:\n")
		board.WriteString(fmt.Sprintf("Turn: %s\n", parts[1]))
		if len(parts) > 2 {
			board.WriteString(fmt.Sprintf("Castling: %s\n", parts[2]))
		}
		if len(parts) > 3 {
			board.WriteString(fmt.Sprintf("En Passant: %s\n", parts[3]))
		}
		if len(parts) > 4 {
			board.WriteString(fmt.Sprintf("Half-move clock: %s\n", parts[4]))
		}
		if len(parts) > 5 {
			board.WriteString(fmt.Sprintf("Full-move number: %s\n", parts[5]))
		}
	}

	return board.String()
}

func main() {
	whitePlayerURL := os.Getenv("WHITE_PLAYER_URL")
	blackPlayerURL := os.Getenv("BLACK_PLAYER_URL")

	coordinator := NewChessCoordinator(whitePlayerURL, blackPlayerURL)

	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req types.MoveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		move, err := coordinator.getMoveFromPlayer(req.FEN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := coordinator.makeMove(move); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(types.MoveResponse{Move: move})
	})

	// Add visualization endpoint
	http.HandleFunc("/visualize", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req types.MoveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		asciiBoard := fenToASCII(req.FEN)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", asciiBoard)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting coordinator on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 