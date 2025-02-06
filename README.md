# Envoy

## Files

### piece.go
- Defines chess piece types and constants
- Provides piece validation and FEN character conversion functions
- Handles piece color and type identification

### board.go
- Implements the chess board structure
- Manages board state and initialization
- Handles piece placement and board operations

### fen.go
- Implements FEN (Forsyth-Edwards Notation) string parsing
- Provides FEN string generation from board state
- Handles board position serialization/deserialization

### move.go
- Defines move structure and validation
- Implements move execution logic
- Handles special moves (castling, en passant)
- Provides move string representation

### game.go
- Manages game state and progression
- Implements game over detection
- Provides game result determination
- Handles win/loss/draw conditions

## Testing

Each component has corresponding test files:
- `piece_test.go`: Tests piece type handling and conversion
- `board_test.go`: Tests board operations and state management
- `fen_test.go`: Tests FEN string parsing and generation
- `move_test.go`: Tests move validation and execution
- `game_test.go`: Tests game state and result determination

## Usage

The package provides a complete chess engine implementation that can be used to:
1. Create and manage chess board positions
2. Validate and execute moves
3. Track game state and determine results
4. Serialize/deserialize board positions using FEN notation

## Dependencies

The package is self-contained and doesn't require external dependencies. 