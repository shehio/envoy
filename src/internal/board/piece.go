package board

// Piece represents a chess piece
type Piece int

const (
	NoPiece Piece = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

// IsWhitePiece returns true if the piece is white
func (p Piece) IsWhitePiece() bool {
	return p > NoPiece && p < BlackPawn
}

// IsEmpty returns true if the square is empty
func (p Piece) IsEmpty() bool {
	return p == NoPiece
}

// isValidPiece returns whether the piece is valid
func isValidPiece(p Piece) bool {
	return p >= WhitePawn && p <= BlackKing
}

// pieceToChar converts a piece to its FEN character representation
func pieceToChar(piece Piece) string {
	switch piece {
	case WhitePawn:
		return "P"
	case WhiteKnight:
		return "N"
	case WhiteBishop:
		return "B"
	case WhiteRook:
		return "R"
	case WhiteQueen:
		return "Q"
	case WhiteKing:
		return "K"
	case BlackPawn:
		return "p"
	case BlackKnight:
		return "n"
	case BlackBishop:
		return "b"
	case BlackRook:
		return "r"
	case BlackQueen:
		return "q"
	case BlackKing:
		return "k"
	default:
		return ""
	}
}

// charToPiece converts a FEN character to a piece
func charToPiece(c rune) Piece {
	switch c {
	case 'P':
		return WhitePawn
	case 'N':
		return WhiteKnight
	case 'B':
		return WhiteBishop
	case 'R':
		return WhiteRook
	case 'Q':
		return WhiteQueen
	case 'K':
		return WhiteKing
	case 'p':
		return BlackPawn
	case 'n':
		return BlackKnight
	case 'b':
		return BlackBishop
	case 'r':
		return BlackRook
	case 'q':
		return BlackQueen
	case 'k':
		return BlackKing
	default:
		return NoPiece
	}
} 