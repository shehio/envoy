package board

import (
	"testing"
)

func TestIsWhitePiece(t *testing.T) {
	tests := []struct {
		piece    Piece
		expected bool
	}{
		{NoPiece, false},
		{WhitePawn, true},
		{WhiteKnight, true},
		{WhiteBishop, true},
		{WhiteRook, true},
		{WhiteQueen, true},
		{WhiteKing, true},
		{BlackPawn, false},
		{BlackKnight, false},
		{BlackBishop, false},
		{BlackRook, false},
		{BlackQueen, false},
		{BlackKing, false},
	}

	for _, test := range tests {
		result := test.piece.IsWhitePiece()
		if result != test.expected {
			t.Errorf("IsWhitePiece(%v) = %v; want %v", test.piece, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		piece    Piece
		expected bool
	}{
		{NoPiece, true},
		{WhitePawn, false},
		{WhiteKnight, false},
		{WhiteBishop, false},
		{WhiteRook, false},
		{WhiteQueen, false},
		{WhiteKing, false},
		{BlackPawn, false},
		{BlackKnight, false},
		{BlackBishop, false},
		{BlackRook, false},
		{BlackQueen, false},
		{BlackKing, false},
	}

	for _, test := range tests {
		result := test.piece.IsEmpty()
		if result != test.expected {
			t.Errorf("IsEmpty(%v) = %v; want %v", test.piece, result, test.expected)
		}
	}
}

func TestIsValidPiece(t *testing.T) {
	tests := []struct {
		piece    Piece
		expected bool
	}{
		{NoPiece, false},
		{WhitePawn, true},
		{WhiteKnight, true},
		{WhiteBishop, true},
		{WhiteRook, true},
		{WhiteQueen, true},
		{WhiteKing, true},
		{BlackPawn, true},
		{BlackKnight, true},
		{BlackBishop, true},
		{BlackRook, true},
		{BlackQueen, true},
		{BlackKing, true},
	}

	for _, test := range tests {
		result := isValidPiece(test.piece)
		if result != test.expected {
			t.Errorf("isValidPiece(%v) = %v; want %v", test.piece, result, test.expected)
		}
	}
}

func TestPieceToChar(t *testing.T) {
	tests := []struct {
		piece    Piece
		expected string
	}{
		{NoPiece, ""},
		{WhitePawn, "P"},
		{WhiteKnight, "N"},
		{WhiteBishop, "B"},
		{WhiteRook, "R"},
		{WhiteQueen, "Q"},
		{WhiteKing, "K"},
		{BlackPawn, "p"},
		{BlackKnight, "n"},
		{BlackBishop, "b"},
		{BlackRook, "r"},
		{BlackQueen, "q"},
		{BlackKing, "k"},
	}

	for _, test := range tests {
		result := pieceToChar(test.piece)
		if result != test.expected {
			t.Errorf("pieceToChar(%v) = %v; want %v", test.piece, result, test.expected)
		}
	}
} 