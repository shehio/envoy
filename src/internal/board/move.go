package board

import (
	"fmt"
	"strings"
)

// Move represents a chess move
type Move struct {
	From      string
	To        string
	Promotion Piece
}

// String returns the move in standard algebraic notation
func (m Move) String() string {
	if m.Promotion != NoPiece {
		return fmt.Sprintf("%s%s%s", m.From, m.To, strings.ToLower(pieceToChar(m.Promotion)))
	}
	return fmt.Sprintf("%s%s", m.From, m.To)
}

// MakeMove makes a move on the board
func (b *Board) MakeMove(move Move) error {
	from := squareToIndex(move.From)
	to := squareToIndex(move.To)
	if from == -1 || to == -1 {
		return fmt.Errorf("invalid square")
	}

	// Handle promotion first
	if move.Promotion != NoPiece {
		toRank := to / 8
		piece := b.squares[from]
		if piece == NoPiece {
			return fmt.Errorf("no piece at source square")
		}
		if piece == WhitePawn && toRank == 7 {
			b.squares[to] = move.Promotion
			b.squares[from] = NoPiece
			b.whiteToMove = !b.whiteToMove
			return nil
		} else if piece == BlackPawn && toRank == 0 {
			b.squares[to] = move.Promotion
			b.squares[from] = NoPiece
			b.whiteToMove = !b.whiteToMove
			return nil
		} else {
			return fmt.Errorf("invalid promotion")
		}
	}

	piece := b.squares[from]
	if piece == NoPiece {
		return fmt.Errorf("no piece at source square")
	}

	// Check if it's the correct player's turn
	if (piece.IsWhitePiece() && !b.whiteToMove) || (!piece.IsWhitePiece() && b.whiteToMove) {
		return fmt.Errorf("not your turn")
	}

	// Handle en passant capture
	if piece == WhitePawn || piece == BlackPawn {
		fromFile := from % 8
		fromRank := from / 8
		toFile := to % 8
		toRank := to / 8

		// Check for en passant capture
		if abs(fromFile-toFile) == 1 && abs(fromRank-toRank) == 1 {
			if b.enPassantSquare != "-" {
				epFile := int(b.enPassantSquare[0] - 'a')
				epRank := int(b.enPassantSquare[1] - '1')
				if toFile == epFile && toRank == epRank {
					// Remove the captured pawn
					capturedRank := fromRank
					if piece == WhitePawn {
						capturedRank = toRank + 1
					} else {
						capturedRank = toRank - 1
					}
					b.squares[capturedRank*8+toFile] = NoPiece
				}
			}
		}

		// Set en passant square for next move
		if abs(fromRank-toRank) == 2 {
			epRank := (fromRank + toRank) / 2
			epFile := fromFile
			b.enPassantSquare = fmt.Sprintf("%c%d", 'a'+epFile, epRank+1)
		} else {
			b.enPassantSquare = "-"
		}
	} else {
		b.enPassantSquare = "-"
	}

	// Make the move
	b.squares[to] = piece
	b.squares[from] = NoPiece

	// Update turn
	b.whiteToMove = !b.whiteToMove

	// Update half move clock
	if piece == WhitePawn || piece == BlackPawn || b.squares[to] != NoPiece {
		b.halfMoveClock = 0
	} else {
		b.halfMoveClock++
	}

	// Update full move number
	if b.whiteToMove {
		b.fullMoveNumber++
	}

	// Update castling rights
	if piece == WhiteKing {
		b.whiteKingsideCastle = false
		b.whiteQueensideCastle = false
	} else if piece == BlackKing {
		b.blackKingsideCastle = false
		b.blackQueensideCastle = false
	} else if piece == WhiteRook {
		if from == 0 {
			b.whiteQueensideCastle = false
		} else if from == 7 {
			b.whiteKingsideCastle = false
		}
	} else if piece == BlackRook {
		if from == 56 {
			b.blackQueensideCastle = false
		} else if from == 63 {
			b.blackKingsideCastle = false
		}
	}

	return nil
}

// Helper functions
func squareToIndex(square string) int {
	if len(square) != 2 {
		return -1
	}
	file := int(square[0] - 'a')
	rank := int(square[1] - '1')
	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return -1
	}
	return rank*8 + file
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
} 