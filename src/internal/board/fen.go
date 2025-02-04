package board

import (
	"fmt"
	"strconv"
	"strings"
)

// SetFEN sets the board position from a FEN string
func (b *Board) SetFEN(fen string) error {
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return fmt.Errorf("invalid FEN string: expected 6 parts, got %d", len(parts))
	}

	// Parse piece placement
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return fmt.Errorf("invalid FEN string: expected 8 ranks, got %d", len(ranks))
	}

	for rank := 7; rank >= 0; rank-- {
		file := 0
		for _, char := range ranks[7-rank] {
			if char >= '1' && char <= '8' {
				emptyCount := int(char - '0')
				for i := 0; i < emptyCount; i++ {
					b.squares[rank*8+file] = NoPiece
					file++
				}
			} else {
				piece := charToPiece(char)
				if piece == NoPiece {
					return fmt.Errorf("invalid piece character: %c", char)
				}
				b.squares[rank*8+file] = piece
				file++
			}
		}
		if file != 8 {
			return fmt.Errorf("invalid rank length: expected 8, got %d", file)
		}
	}

	// Parse side to move
	switch parts[1] {
	case "w":
		b.whiteToMove = true
	case "b":
		b.whiteToMove = false
	default:
		return fmt.Errorf("invalid side to move: %s", parts[1])
	}

	// Parse castling rights
	b.whiteKingsideCastle = false
	b.whiteQueensideCastle = false
	b.blackKingsideCastle = false
	b.blackQueensideCastle = false
	for _, char := range parts[2] {
		switch char {
		case 'K':
			b.whiteKingsideCastle = true
		case 'Q':
			b.whiteQueensideCastle = true
		case 'k':
			b.blackKingsideCastle = true
		case 'q':
			b.blackQueensideCastle = true
		case '-':
			// No castling rights
		default:
			return fmt.Errorf("invalid castling character: %c", char)
		}
	}

	// Parse en passant square
	b.enPassantSquare = parts[3]
	if b.enPassantSquare != "-" {
		if len(b.enPassantSquare) != 2 {
			return fmt.Errorf("invalid en passant square: %s", b.enPassantSquare)
		}
		file := int(b.enPassantSquare[0] - 'a')
		rank := int(b.enPassantSquare[1] - '1')
		if file < 0 || file > 7 || rank < 0 || rank > 7 {
			return fmt.Errorf("invalid en passant square coordinates: %s", b.enPassantSquare)
		}
	}

	// Parse halfmove clock
	halfMoveClock, err := strconv.Atoi(parts[4])
	if err != nil || halfMoveClock < 0 {
		return fmt.Errorf("invalid half move clock")
	}
	b.halfMoveClock = halfMoveClock

	// Parse full move number
	fullMoveNumber, err := strconv.Atoi(parts[5])
	if err != nil || fullMoveNumber < 1 {
		return fmt.Errorf("invalid full move number")
	}
	b.fullMoveNumber = fullMoveNumber

	return nil
}

// FEN returns the current position in FEN notation
func (b *Board) FEN() string {
	var fen strings.Builder

	// Piece placement
	for rank := 7; rank >= 0; rank-- {
		emptyCount := 0
		for file := 0; file < 8; file++ {
			piece := b.squares[rank*8+file]
			if piece == NoPiece {
				emptyCount++
			} else {
				if emptyCount > 0 {
					fen.WriteString(fmt.Sprintf("%d", emptyCount))
					emptyCount = 0
				}
				fen.WriteString(pieceToChar(piece))
			}
		}
		if emptyCount > 0 {
			fen.WriteString(fmt.Sprintf("%d", emptyCount))
		}
		if rank > 0 {
			fen.WriteString("/")
		}
	}

	// Side to move
	if b.whiteToMove {
		fen.WriteString(" w ")
	} else {
		fen.WriteString(" b ")
	}

	// Castling rights
	castling := ""
	if b.whiteKingsideCastle {
		castling += "K"
	}
	if b.whiteQueensideCastle {
		castling += "Q"
	}
	if b.blackKingsideCastle {
		castling += "k"
	}
	if b.blackQueensideCastle {
		castling += "q"
	}
	if castling == "" {
		castling = "-"
	}
	fen.WriteString(castling + " ")

	// En passant square
	fen.WriteString(b.enPassantSquare + " ")

	// Halfmove clock and full move number
	fen.WriteString(fmt.Sprintf("%d %d", b.halfMoveClock, b.fullMoveNumber))

	return fen.String()
}

