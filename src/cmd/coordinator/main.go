package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func (c *ChessCoordinator) makeMove(move string) error {
	return c.board.MakeMove(move)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting coordinator on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 