package board

type Board struct {
	squares [64]Piece
	whiteToMove bool
	whiteKingsideCastle bool
	whiteQueensideCastle bool
	blackKingsideCastle bool
	blackQueensideCastle bool
	halfMoveClock int
	fullMoveNumber int
	enPassantSquare string // The square where en passant capture is possible
}

func NewBoard() *Board {
	board := &Board{
		whiteToMove: true,
		whiteKingsideCastle: true,
		whiteQueensideCastle: true,
		blackKingsideCastle: true,
		blackQueensideCastle: true,
		fullMoveNumber: 1,
		enPassantSquare: "-",
	}

	initialPosition := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board.SetFEN(initialPosition)

	return board
}

func (b *Board) GetFEN() string {
	return b.FEN()
}

// IsWhiteToMove returns whether it's white's turn to move
func (b *Board) IsWhiteToMove() bool {
	return b.whiteToMove
}

// IsWhiteTurn returns true if it's white's turn to move
func (b *Board) IsWhiteTurn() bool {
	return b.whiteToMove
}

func (b *Board) SetEnPassantSquare(square string) {
	b.enPassantSquare = square
}

func (b *Board) GetEnPassantSquare() string {
	return b.enPassantSquare
} 