package stockfish

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Handler manages communication with the Stockfish chess engine
type Handler struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

// NewHandler creates a new Stockfish handler
func NewHandler() (*Handler, error) {
	cmd := exec.Command("stockfish")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start stockfish: %v", err)
	}

	handler := &Handler{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
	}

	if err := handler.initializeEngine(); err != nil {
		return nil, err
	}

	return handler, nil
}

func (h *Handler) initializeEngine() error {
	// Send UCI command
	if _, err := fmt.Fprintln(h.stdin, "uci"); err != nil {
		return fmt.Errorf("failed to send uci command: %v", err)
	}

	// Wait for uciok
	scanner := bufio.NewScanner(h.stdout)
	for scanner.Scan() {
		if scanner.Text() == "uciok" {
			break
		}
	}

	// Send isready command
	if _, err := fmt.Fprintln(h.stdin, "isready"); err != nil {
		return fmt.Errorf("failed to send isready command: %v", err)
	}

	// Wait for readyok
	for scanner.Scan() {
		if scanner.Text() == "readyok" {
			break
		}
	}

	return nil
}

// GetMove implements the Player interface
func (h *Handler) GetMove(fen string) (string, error) {
	// Set position
	if _, err := fmt.Fprintf(h.stdin, "position fen %s\n", fen); err != nil {
		return "", fmt.Errorf("failed to set position: %v", err)
	}

	// Start thinking
	if _, err := fmt.Fprintf(h.stdin, "go movetime %d\n", 1000); err != nil {
		return "", fmt.Errorf("failed to start thinking: %v", err)
	}

	// Read best move
	scanner := bufio.NewScanner(h.stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "bestmove") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	return "", fmt.Errorf("failed to get best move")
}

// Close closes the Stockfish process
func (h *Handler) Close() error {
	if h.cmd.Process != nil {
		return h.cmd.Process.Kill()
	}
	return nil
} 