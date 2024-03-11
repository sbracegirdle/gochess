# Gochess

This is a simple implementation of a chess game in Go for learning purposes. It includes the basic rules of chess, including piece movement and game state management.

## Files

`game.go`
This file contains the main game logic. It includes the following:

`NewGame(player1Name, player2Name string) *Game`: This function initializes a new game with two players. It sets up the board and the pieces for each player.

`MovePiece(currentX, currentY, newX, newY int) error`: This function moves a piece from one position to another. It checks if the move is valid and updates the game state accordingly.

`moves.go`

This file contains the logic for validating the moves of each piece. It includes the following:

`IsValidMove(color PieceColor, currentX, currentY, newX, newY int) error`: This function checks if a move is valid for a given piece.
