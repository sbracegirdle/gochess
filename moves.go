package main

import (
	"errors"
)

func (g *Game) IsValidMove(color PieceColor, currentX, currentY, newX, newY int) error {
	// Check if the current position has a piece
	if g.Board[currentX][currentY] == nil {
		return errors.New("no piece at the current position")
	}

	// Check if the piece at the current position is of the same color as the player
	if g.Board[currentX][currentY].Color != color {
		return errors.New("piece at the current position is not of the player's color")
	}

	// Check if the new position is within the board
	if newX < 0 || newX > 7 || newY < 0 || newY > 7 {
		return errors.New("new position is out of bounds")
	}

	if g.Board[newX][newY] != nil {
		if g.Board[newX][newY].Color == g.Board[currentX][currentY].Color {
			return errors.New("cannot capture your own piece")
		}
	}

	if g.WouldBeCheck(color, currentX, currentY, newX, newY) {
		return errors.New("move would result in check")
	}

	piece := g.Board[currentX][currentY]

	switch piece.Type {
	case Pawn:
		return g.IsValidPawnMove(currentX, currentY, newX, newY)
	case Rook:
		return g.IsValidRookMove(currentX, currentY, newX, newY)
	case Knight:
		return g.IsValidKnightMove(currentX, currentY, newX, newY)
	case Bishop:
		return g.IsValidBishopMove(currentX, currentY, newX, newY)
	case Queen:
		return g.IsValidQueenMove(currentX, currentY, newX, newY)
	case King:
		return g.IsValidKingMove(currentX, currentY, newX, newY)
	default:
		return errors.New("Unknown piece type found: " + string(piece.Type))
	}
}

func (g *Game) IsValidPawnMove(currentX, currentY, newX, newY int) error {
	positiveYDirection := 1
	startingRow := 1
	oppositeColor := Black
	if g.Board[currentX][currentY].Color == Black {
		positiveYDirection = -1
		startingRow = 6
		oppositeColor = White
	}

	// Check if target position & position in between are unoccupied
	isRoadClear := g.IsPathClear(currentX, currentY, newX, newY, true)

	if isRoadClear && newX == currentX && newY == currentY+positiveYDirection {
		// Can move one step forward if the new position is empty
		return nil
	} else if isRoadClear && newX == currentX && currentY == startingRow && newY == currentY+2*positiveYDirection && g.Board[newX][newY-positiveYDirection] == nil {
		// Can move two steps forward if it's the pawn's first move (in starting row) and the new position is empty and the position in between is empty
		return nil
	} else if (newX == currentX-1 || newX == currentX+1) && (newY == currentY+positiveYDirection) && g.Board[newX][newY].Color == oppositeColor {
		// Can capture a piece if it's one step diagonally forward and the new position has a piece of the opposite color
		return nil
	} else {
		return errors.New("invalid move for white pawn")
	}
}

func (g *Game) IsValidKnightMove(currentX, currentY, newX, newY int) error {
	// Check if the new position is within the board
	if newX < 0 || newX > 7 || newY < 0 || newY > 7 {
		return errors.New("move is out of board")
	}

	// Check if the new position is two squares horizontally and one square vertically
	// or two squares vertically and one square horizontally from the current position
	if (abs(newX-currentX) == 2 && abs(newY-currentY) == 1) || (abs(newX-currentX) == 1 && abs(newY-currentY) == 2) {
		// Check if the new position is empty or has a piece of the opposite color
		if g.Board[newX][newY] == nil || g.Board[newX][newY].Color != g.Board[currentX][currentY].Color {
			return nil
		} else {
			return errors.New("new position has a piece of the same color")
		}
	} else {
		return errors.New("invalid move for knight")
	}
}

func (g *Game) IsValidBishopMove(currentX, currentY, newX, newY int) error {
	dx := abs(newX - currentX)
	dy := abs(newY - currentY)

	// Check if the move is diagonal
	if dx != dy {
		return errors.New("invalid move for bishop: not diagonal")
	}

	// Check if the path is clear
	if !g.IsPathClear(currentX, currentY, newX, newY, false) {
		return errors.New("invalid move for bishop: path is not clear")
	}

	// If the target position is occupied by a piece of the same color
	if g.Board[newX][newY] != nil && g.Board[newX][newY].Color == g.Board[currentX][currentY].Color {
		return errors.New("invalid move for bishop: cannot capture own piece")
	}

	return nil
}

func (g *Game) IsValidRookMove(currentX, currentY, newX, newY int) error {
	dx := sign(newX - currentX)
	dy := sign(newY - currentY)

	isRankFile := dx != dy && (dx == 0 || dy == 0)

	// Check if the move is along a rank or file
	if !isRankFile {
		return errors.New("invalid move for rook: not along a rank or file")
	}

	// Check if the path is clear
	if !g.IsPathClear(currentX, currentY, newX, newY, false) {
		return errors.New("invalid move for rook: path is not clear")
	}

	// If the target position is occupied by a piece of the same color
	if g.Board[newX][newY] != nil && g.Board[newX][newY].Color == g.Board[currentX][currentY].Color {
		return errors.New("invalid move for rook: cannot capture own piece")
	}

	return nil
}

func (g *Game) IsValidQueenMove(currentX, currentY, newX, newY int) error {
	if err := g.IsValidRookMove(currentX, currentY, newX, newY); err == nil {
		return nil
	}

	if err := g.IsValidBishopMove(currentX, currentY, newX, newY); err == nil {
		return nil
	}

	return errors.New("invalid move for queen")
}

func (g *Game) IsValidKingMove(currentX, currentY, newX, newY int) error {
	// If the target position is occupied by a piece of the same color
	if g.Board[newX][newY] != nil && g.Board[newX][newY].Color == g.Board[currentX][currentY].Color {
		return errors.New("invalid move for king: cannot capture own piece")
	}

	if abs(newX-currentX) > 1 || abs(newY-currentY) > 1 {
		return errors.New("invalid move for king")
	}

	return nil
}

func (g *Game) IsPathClear(startX, startY, endX, endY int, includeEnd bool) bool {
	dx := endX - startX
	dy := endY - startY

	stepX, stepY := 0, 0
	if dx != 0 {
		stepX = dx / abs(dx)
	}
	if dy != 0 {
		stepY = dy / abs(dy)
	}

	x, y := startX+stepX, startY+stepY
	for (x != endX || y != endY) && x >= 0 && x < 8 && y >= 0 && y < 8 {
		if g.Board[x][y] != nil {
			return false
		}
		x += stepX
		y += stepY
	}

	// If includeEnd is true, check the end position as well
	if includeEnd && g.Board[endX][endY] != nil {
		return false
	}

	return true
}

func (g *Game) FindKing(color PieceColor) (int, int) {
	var kingX, kingY int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g.Board[i][j] != nil && g.Board[i][j].Type == King && g.Board[i][j].Color == color {
				kingX, kingY = i, j
			}
		}
	}
	return kingX, kingY
}

func (g *Game) IsCheck(color PieceColor) bool {
	// Find the king
	kingX, kingY := g.FindKing(color)

	// Check if any piece of the opposite color can move to the king's position
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g.Board[i][j] != nil && g.Board[i][j].Color != color {
				if g.IsValidMove(g.Board[i][j].Color, i, j, kingX, kingY) == nil {
					return true
				}
			}
		}
	}

	return false
}

func (g *Game) WouldBeCheck(color PieceColor, currentX, currentY, newX, newY int) bool {
	// Save the current state of the board and the game history
	savedBoard := g.Board
	savedHistory := g.History

	// Perform the move
	g.Board[newX][newY] = g.Board[currentX][currentY]
	g.Board[currentX][currentY] = nil

	// Check if the move results in a check
	isCheck := g.IsCheck(color)

	// Revert the move by restoring the saved state of the board and the game history
	g.Board = savedBoard
	g.History = savedHistory

	// Return the result of the check
	return isCheck
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	return x / abs(x)
}
