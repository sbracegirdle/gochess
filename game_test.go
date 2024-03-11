package main

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame("Alice", "Bob")

	// Check that the game is in the ongoing state
	if game.State != Ongoing {
		t.Errorf("Expected game state to be %v, but got %v", Ongoing, game.State)
	}

	// Check that the players are set up correctly
	if game.Players[0].Name != "Alice" || game.Players[0].Color != White {
		t.Errorf("Expected first player to be Alice (White), but got %v (%v)", game.Players[0].Name, game.Players[0].Color)
	}
	if game.Players[1].Name != "Bob" || game.Players[1].Color != Black {
		t.Errorf("Expected second player to be Bob (Black), but got %v (%v)", game.Players[1].Name, game.Players[1].Color)
	}

	// Check that all pieces are set up correctly
	for i := 0; i < 8; i++ {
		if game.Board[1][i] == nil || game.Board[1][i].Type != Pawn || game.Board[1][i].Color != Black {
			t.Errorf("Expected a black pawn at position (1, %d), but got %v", i, game.Board[1][i])
		}
		if game.Board[6][i] == nil || game.Board[6][i].Type != Pawn || game.Board[6][i].Color != White {
			t.Errorf("Expected a white pawn at position (6, %d), but got %v", i, game.Board[6][i])
		}
	}

	pieceTypes := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for i, pieceType := range pieceTypes {
		if game.Board[0][i] == nil || game.Board[0][i].Type != pieceType || game.Board[0][i].Color != Black {
			t.Errorf("Expected a black %v at position (0, %d), but got %v", pieceType, i, game.Board[0][i])
		}
		if game.Board[7][i] == nil || game.Board[7][i].Type != pieceType || game.Board[7][i].Color != White {
			t.Errorf("Expected a white %v at position (7, %d), but got %v", pieceType, i, game.Board[7][i])
		}
	}

}


func TestIsCheckmate(t *testing.T) {
	tests := []struct {
		name string
		color PieceColor
		board Board
		want bool
	}{
		{
			name: "White is in checkmate",
			color: White,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{0, 4}: {Color: White, Type: King},
				{0, 2}: {Color: Black, Type: Queen},
				{1, 6}: {Color: Black, Type: Rook},
			}),
			want: true,
		},
		{
			name: "Black is not in checkmate",
			color: Black,
			board: createBoardWithPieces(map[[2]int]*Piece{
				{7, 4}: {Color: Black, Type: King},
				{6, 3}: {Color: White, Type: Pawn},
			}),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{Board: tt.board}
			if got := g.IsCheckmate(tt.color); got != tt.want {
				t.Errorf("IsCheckmate() = %v, want %v", got, tt.want)
			}
		})
	}
}