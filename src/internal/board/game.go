package board

// IsGameOver returns whether the game is over
func (b *Board) IsGameOver() bool {
	whiteKingFound := false
	blackKingFound := false

	for _, piece := range b.squares {
		if piece == WhiteKing {
			whiteKingFound = true
		} else if piece == BlackKing {
			blackKingFound = true
		}
	}

	return !whiteKingFound || !blackKingFound
}

// Result returns the game result ("1-0", "0-1", "1/2-1/2", or "*")
func (b *Board) Result() string {
	if !b.IsGameOver() {
		return "*"
	}

	whiteKingFound := false
	blackKingFound := false

	for _, piece := range b.squares {
		if piece == WhiteKing {
			whiteKingFound = true
		} else if piece == BlackKing {
			blackKingFound = true
		}
	}

	if !whiteKingFound && !blackKingFound {
		return "1/2-1/2"
	}
	if !whiteKingFound {
		return "0-1"
	}
	if !blackKingFound {
		return "1-0"
	}

	return "1/2-1/2"
} 