package round

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
)

// Status is the current state of the round
type Status int

// InProgress is the state when the round has no winners
const InProgress Status = 0

// WinnerSnake1 is the state when snake 1 has won the round
const WinnerSnake1 Status = 1

// WinnerSnake2 is the state when snake 2 has won the round
const WinnerSnake2 Status = 2

// Round simulates a round
type Round struct {
	status     Status
	playerTurn bool
	player1    player.Player
	player2    player.Player
	gridColor  sdl.Color
}
