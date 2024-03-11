package main

import "errors"

type PieceType string

const (
	Pawn   PieceType = "Pawn"
	Rook   PieceType = "Rook"
	Knight PieceType = "Knight"
	Bishop PieceType = "Bishop"
	Queen  PieceType = "Queen"
	King   PieceType = "King"
)

type PieceColor string

const (
	White PieceColor = "White"
	Black PieceColor = "Black"
)

type Piece struct {
	Type  PieceType
	Color PieceColor
}

type Board [8][8]*Piece

type Player struct {
	Name  string
	Color PieceColor
}

type Position struct {
	X int
	Y int
}

type Move struct {
	Color      PieceColor
	From       Position
	To         Position
	PieceTaken *Piece
}

type GameState string

const (
	Ongoing      GameState = "Ongoing"
	WhiteWon     GameState = "WhiteWon"
	BlackWon     GameState = "BlackWon"
	Draw         GameState = "Draw"
	PromoteWhite GameState = "PromoteWhite"
	PromoteBlack GameState = "PromoteBlack"
)

type Game struct {
	Board      Board
	Players    [2]Player
	State      GameState
	PlayerTurn PieceColor
	History    []Move
}

func NewGame(player1Name, player2Name string) *Game {
	// Initialize an empty board
	var board Board

	// Set up pawns
	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{Type: Pawn, Color: Black}
		board[6][i] = &Piece{Type: Pawn, Color: White}
	}

	// Set up rooks
	board[0][0] = &Piece{Type: Rook, Color: Black}
	board[0][7] = &Piece{Type: Rook, Color: Black}
	board[7][0] = &Piece{Type: Rook, Color: White}
	board[7][7] = &Piece{Type: Rook, Color: White}

	// Set up knights
	board[0][1] = &Piece{Type: Knight, Color: Black}
	board[0][6] = &Piece{Type: Knight, Color: Black}
	board[7][1] = &Piece{Type: Knight, Color: White}
	board[7][6] = &Piece{Type: Knight, Color: White}

	// Set up bishops
	board[0][2] = &Piece{Type: Bishop, Color: Black}
	board[0][5] = &Piece{Type: Bishop, Color: Black}
	board[7][2] = &Piece{Type: Bishop, Color: White}
	board[7][5] = &Piece{Type: Bishop, Color: White}

	// Set up queens
	board[0][3] = &Piece{Type: Queen, Color: Black}
	board[7][3] = &Piece{Type: Queen, Color: White}

	// Set up kings
	board[0][4] = &Piece{Type: King, Color: Black}
	board[7][4] = &Piece{Type: King, Color: White}

	// Create the players
	player1 := Player{Name: player1Name, Color: White}
	player2 := Player{Name: player2Name, Color: Black}

	// Create the game
	game := Game{
		Board:   board,
		Players: [2]Player{player1, player2},
		State:   Ongoing,
	}

	return &game
}

func (g *Game) MovePiece(currentX, currentY, newX, newY int) error {
	// Check if game state is valid
	if g.State != Ongoing {
		return errors.New("Game is not ongoing, got state: " + string(g.State))
	}

	// Infer the current player's color from the game history. If no history, assume white
	currentPlayerColor := White
	if len(g.History) > 0 {
		// Need to invert the color of the last move
		if g.History[len(g.History)-1].Color == White {
			currentPlayerColor = Black
		}
	}
	otherPlayerColor := Black
	if currentPlayerColor == Black {
		otherPlayerColor = White
	}

	err := g.IsValidMove(currentPlayerColor, currentX, currentY, newX, newY)

	if err != nil {
		return err
	}

	// Move the piece
	// See if piece is being taken
	pieceTaken := g.Board[newX][newY]
	g.Board[newX][newY] = g.Board[currentX][currentY]
	g.Board[currentX][currentY] = nil
	g.History = append(g.History, Move{
		Color:      currentPlayerColor,
		From:       Position{X: currentX, Y: currentY},
		To:         Position{X: newX, Y: newY},
		PieceTaken: pieceTaken,
	})

	// TODO Special moves like castling, en passant, pawn promotion, etc.

	// Check if the game is over
	if g.IsCheckmate(otherPlayerColor) {
		g.State = WhiteWon
		if otherPlayerColor == White {
			g.State = BlackWon
		}
	}
	// TODO Check for draw

	return nil
}

func (g *Game) IsCheckmate(color PieceColor) bool {
	// Check if the king is in check
	if !g.IsCheck(color) {
		return false
	}

	// Find the king
	kingX, kingY := g.FindKing(color)

	// Check if the king has any valid moves
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				newX, newY := kingX+dx, kingY+dy
				if newX >= 0 && newX < 8 && newY >= 0 && newY < 8 && g.IsValidMove(color, kingX, kingY, newX, newY) == nil {
					return false
				}
			}
		}
	}

	return true
}
