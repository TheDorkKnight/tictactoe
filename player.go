package main

type Player interface {
	NextBoard(b Board) Board
}

type PlayerID int
const (
	NoPlayer PlayerID = iota
	PlayerOne
	PlayerTwo
)

func (pid PlayerID) String() string {
	switch (pid) {
	case PlayerOne:
		return "X"
	case PlayerTwo:
		return "O"
	}
	return " "
}

func OtherPlayerID(pid PlayerID) PlayerID {
	if pid == PlayerOne {
		return PlayerTwo
	}
	return PlayerOne
}

type PlayerKind int
const (
	AIPlayer PlayerKind = iota
	HumanPlayer
)
