package main

import (
	"errors"
	"testing"
)

func TestIsValidMove(t *testing.T) {
	tests := []struct {
		name    string
		startX  int
		startY  int
		endX    int
		endY    int
		board   Board
		wantErr bool
	}{
		{
			name:    "no piece at the current position",
			startX:  3,
			startY:  3,
			endX:    4,
			endY:    4,
			board:   createEmptyBoard(),
			wantErr: true,
		},
		{
			name:    "new position is out of bounds",
			startX:  3,
			startY:  3,
			endX:    8,
			endY:    8,
			board:   createBoardWithPieces(map[[2]int]*Piece{{3, 3}: {Color: "White", Type: "Pawn"}}),
			wantErr: true,
		},
		{
			name:   "cannot capture your own piece",
			startX: 3,
			startY: 3,
			endX:   4,
			endY:   4,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{3, 3}: {Color: "White", Type: "Pawn"},
				{4, 4}: {Color: "White", Type: "Pawn"},
			}),
			wantErr: true,
		},
		{
			name:   "valid move",
			startX: 3,
			startY: 3,
			endX:   4,
			endY:   4,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{3, 3}: {Color: "White", Type: "Pawn"},
				{4, 4}: {Color: "Black", Type: "Pawn"},
			}),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}

			if err := g.IsValidMove("White", tt.startX, tt.startY, tt.endX, tt.endY); (err != nil) != tt.wantErr {
				t.Errorf("IsValidMove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidPawnMove(t *testing.T) {
	tests := []struct {
		name          string
		currentX      int
		currentY      int
		newX          int
		newY          int
		board         [8][8]*Piece
		expectedError error
	}{
		{
			name:     "valid one step forward move for white pawn",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     2,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: nil,
		},
		{
			name:     "blocked one step forward move for white pawn",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     2,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: errors.New("invalid move for white pawn"),
		},
		{
			name:     "valid two steps forward move for white pawn on starting row",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     3,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: nil,
		},
		{
			name:     "blocked two steps forward move for white pawn on starting row",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     3,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: errors.New("invalid move for white pawn"),
		},
		{
			name:     "valid capture move for white pawn",
			currentX: 1,
			currentY: 1,
			newX:     2,
			newY:     2,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil, nil},
				{nil, nil, &Piece{Type: "Pawn", Color: "Black"}, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: nil,
		},
		{
			name:     "invalid capture move for white pawn",
			currentX: 1,
			currentY: 1,
			newX:     2,
			newY:     2,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil, nil},
				{nil, nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: errors.New("invalid move for white pawn"),
		},
		{
			name:     "invalid forward move for white pawn",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     4,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, &Piece{Type: "Pawn", Color: "White"}, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: errors.New("invalid move for white pawn"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Board: tt.board,
			}
			err := g.IsValidPawnMove(tt.currentX, tt.currentY, tt.newX, tt.newY)
			if err != nil && tt.expectedError != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("got %v, want %v", err, tt.expectedError)
				}
			} else if err != tt.expectedError {
				t.Errorf("got %v, want %v", err, tt.expectedError)
			}
		})
	}
}

func TestIsValidKnightMove(t *testing.T) {
	tests := []struct {
		name          string
		currentX      int
		currentY      int
		newX          int
		newY          int
		board         [8][8]*Piece
		expectedError error
	}{
		{
			name:     "valid knight move up-right",
			currentX: 4,
			currentY: 4,
			newX:     5,
			newY:     6,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, &Piece{Type: "Knight", Color: "White"}, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: nil,
		},
		{
			name:     "valid knight move down-left",
			currentX: 4,
			currentY: 4,
			newX:     3,
			newY:     2,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, &Piece{Type: "Knight", Color: "White"}, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: nil,
		},
		{
			name:     "invalid knight move",
			currentX: 4,
			currentY: 4,
			newX:     5,
			newY:     5,
			board: [8][8]*Piece{
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, &Piece{Type: "Knight", Color: "White"}, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
			},
			expectedError: errors.New("invalid move for knight"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Board: tt.board,
			}
			err := g.IsValidKnightMove(tt.currentX, tt.currentY, tt.newX, tt.newY)
			if err != nil && tt.expectedError != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("got %v, want %v", err, tt.expectedError)
				}
			} else if err != tt.expectedError {
				t.Errorf("got %v, want %v", err, tt.expectedError)
			}
		})
	}
}

func TestIsValidBishopMove(t *testing.T) {
	tests := []struct {
		name    string
		startX  int
		startY  int
		endX    int
		endY    int
		board   Board
		wantErr bool
	}{
		{
			name:    "valid diagonal move",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    5,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid diagonal move reverse",
			startX:  5,
			startY:  5,
			endX:    3,
			endY:    3,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "invalid non-diagonal move",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    4,
			board:   createEmptyBoard(),
			wantErr: true,
		},
		{
			name:    "invalid move, path not clear",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    5,
			board:   createBoardWithPieceAt(4, 4, &Piece{Color: "white", Type: "pawn"}),
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			g.Board[tt.startX][tt.startY] = &Piece{Color: "white", Type: "bishop"}

			if err := g.IsValidBishopMove(tt.startX, tt.startY, tt.endX, tt.endY); (err != nil) != tt.wantErr {
				t.Errorf("IsValidBishopMove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidRookMove(t *testing.T) {
	tests := []struct {
		name    string
		startX  int
		startY  int
		endX    int
		endY    int
		board   Board
		wantErr bool
	}{
		{
			name:    "valid horizontal move",
			startX:  3,
			startY:  3,
			endX:    3,
			endY:    5,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid vertical move",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    3,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "invalid diagonal move",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    5,
			board:   createEmptyBoard(),
			wantErr: true,
		},
		{
			name:    "invalid move, path not clear",
			startX:  3,
			startY:  3,
			endX:    3,
			endY:    5,
			board:   createBoardWithPieceAt(3, 4, &Piece{Color: "White", Type: "Pawn"}),
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			g.Board[tt.startX][tt.startY] = &Piece{Color: "white", Type: "rook"}

			if err := g.IsValidRookMove(tt.startX, tt.startY, tt.endX, tt.endY); (err != nil) != tt.wantErr {
				t.Errorf("IsValidRookMove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidQueenMove(t *testing.T) {
	tests := []struct {
		name     string
		currentX int
		currentY int
		newX     int
		newY     int
		board    Board
		wantErr  bool
	}{
		{
			name:     "valid horizontal move",
			currentX: 1,
			currentY: 1,
			newX:     1,
			newY:     5,
			board:    createEmptyBoard(),
			wantErr:  false,
		},
		{
			name:     "valid vertical move",
			currentX: 1,
			currentY: 1,
			newX:     5,
			newY:     1,
			board:    createEmptyBoard(),
			wantErr:  false,
		},
		{
			name:     "valid diagonal move",
			currentX: 1,
			currentY: 1,
			newX:     3,
			newY:     3,
			board:    createEmptyBoard(),
			wantErr:  false,
		},
		{
			name:     "invalid move",
			currentX: 1,
			currentY: 1,
			newX:     2,
			newY:     3,
			board:    createEmptyBoard(),
			wantErr:  true,
		},
		{
			name:     "invalid move, path not clear",
			currentX: 3,
			currentY: 3,
			newX:     3,
			newY:     5,
			board:    createBoardWithPieceAt(3, 4, &Piece{Color: "White", Type: "Pawn"}),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			if err := g.IsValidQueenMove(tt.currentX, tt.currentY, tt.newX, tt.newY); (err != nil) != tt.wantErr {
				t.Errorf("IsValidQueenMove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidKingMove(t *testing.T) {
	tests := []struct {
		name    string
		startX  int
		startY  int
		endX    int
		endY    int
		board   Board
		wantErr bool
	}{
		{
			name:    "valid move up",
			startX:  3,
			startY:  3,
			endX:    3,
			endY:    4,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid move down",
			startX:  3,
			startY:  3,
			endX:    3,
			endY:    2,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid move left",
			startX:  3,
			startY:  3,
			endX:    2,
			endY:    3,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid move right",
			startX:  3,
			startY:  3,
			endX:    4,
			endY:    3,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "valid move diagonal",
			startX:  3,
			startY:  3,
			endX:    4,
			endY:    4,
			board:   createEmptyBoard(),
			wantErr: false,
		},
		{
			name:    "invalid move",
			startX:  3,
			startY:  3,
			endX:    5,
			endY:    5,
			board:   createEmptyBoard(),
			wantErr: true,
		},
		{
			name:    "invalid move, path not clear",
			startX:  3,
			startY:  3,
			endX:    4,
			endY:    4,
			board:   createBoardWithPieceAt(4, 4, &Piece{Color: "White", Type: "Pawn"}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			g.Board[tt.startX][tt.startY] = &Piece{Color: "White", Type: "King"}

			if err := g.IsValidKingMove(tt.startX, tt.startY, tt.endX, tt.endY); (err != nil) != tt.wantErr {
				t.Errorf("IsValidKingMove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsPathClear(t *testing.T) {
	tests := []struct {
		name       string
		board      Board
		startX     int
		startY     int
		endX       int
		endY       int
		includeEnd bool
		expected   bool
	}{
		{
			name:       "Clear path",
			board:      createEmptyBoard(),
			startX:     0,
			startY:     0,
			endX:       7,
			endY:       7,
			includeEnd: false,
			expected:   true,
		},
		{
			name:       "Obstructed path",
			board:      createBoardWithPieceAt(3, 3, &Piece{Color: "white", Type: "pawn"}),
			startX:     0,
			startY:     0,
			endX:       7,
			endY:       7,
			includeEnd: false,
			expected:   false,
		},
		{
			name:       "Clear path one space",
			board:      createEmptyBoard(),
			startX:     0,
			startY:     0,
			endX:       1,
			endY:       0,
			includeEnd: false,
			expected:   true,
		},
		{
			name:       "Piece at destination is fine",
			board:      createBoardWithPieceAt(1, 0, &Piece{Color: "white", Type: "pawn"}),
			startX:     0,
			startY:     0,
			endX:       1,
			endY:       0,
			includeEnd: false,
			expected:   true,
		},
		{
			name:       "Piece at destination is not fine if including",
			board:      createBoardWithPieceAt(1, 0, &Piece{Color: "white", Type: "pawn"}),
			startX:     0,
			startY:     0,
			endX:       1,
			endY:       0,
			includeEnd: true,
			expected:   false,
		},
		{
			name:       "Piece at start is fine",
			board:      createBoardWithPieceAt(0, 0, &Piece{Color: "white", Type: "pawn"}),
			startX:     0,
			startY:     0,
			endX:       1,
			endY:       0,
			includeEnd: false,
			expected:   true,
		},
		{
			name:       "Empty move is fine",
			board:      createBoardWithPieceAt(0, 0, &Piece{Color: "white", Type: "pawn"}),
			startX:     0,
			startY:     0,
			endX:       0,
			endY:       0,
			includeEnd: false,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Board: tt.board,
			}
			if got := g.IsPathClear(tt.startX, tt.startY, tt.endX, tt.endY, tt.includeEnd); got != tt.expected {
				t.Errorf("IsPathClear() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func createEmptyBoard() Board {
	// board is [8][8]*Piece
	return [8][8]*Piece{
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
	}
}

func createBoardWithPieceAt(x, y int, piece *Piece) Board {
	board := createEmptyBoard()
	board[x][y] = piece
	return board
}

func addPiece(board Board, x, y int, piece *Piece) Board {
	board[x][y] = piece
	return board
}

// Give list of pieces to add to the board, and their positions
func createBoardWithPieces(pieces map[[2]int]*Piece) Board {
	board := createEmptyBoard()
	for position, piece := range pieces {
		board = addPiece(board, position[0], position[1], piece)
	}
	return board
}

func TestIsCheck(t *testing.T) {
	tests := []struct {
		name  string
		color PieceColor
		board Board
		want  bool
	}{
		{
			name:  "White is in check",
			color: White,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{0, 4}: {Color: White, Type: King},
				{1, 5}: {Color: Black, Type: Queen},
			}),
			want: true,
		},
		{
			name:  "Black is not in check",
			color: Black,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{7, 4}: {Color: Black, Type: King},
				{1, 1}: {Color: White, Type: Pawn},
			}),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			if got := g.IsCheck(tt.color); got != tt.want {
				t.Errorf("IsCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
