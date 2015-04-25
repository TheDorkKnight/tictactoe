package main

type Game interface {
	String() string
	CurrentPlayer() PlayerID
	Next() Game
	IsOver() (bool, PlayerID)
}

type game struct {
	b Board
	current PlayerID
	p1 Player
	p2 Player
}

func MakeGame(p1k PlayerKind, p2k PlayerKind) Game {
	if p1k != AIPlayer && p1k != HumanPlayer {
		p1k = AIPlayer
	}
	if p2k != AIPlayer && p2k != HumanPlayer {
		p2k = AIPlayer
	}

	var p1 Player
	if p1k == AIPlayer {
		p1 = AI{self: PlayerOne}
	} else {
		p1 = Human{self: PlayerOne}
	}

	var p2 Player
	if p2k == AIPlayer {
		p2 = AI{self: PlayerTwo}
	} else {
		p2 = Human{self: PlayerTwo}
	}

	return game{b: &board{}, current: PlayerOne, p1: p1, p2: p2}
}

func (g game) CurrentPlayer() PlayerID {
	return g.current
}

func (g game) String() string {
	return g.b.PrettyString()
}

func (g game) Next() Game {	
	var p Player
	if g.current != PlayerTwo {
		p = g.p1
	} else {
		p = g.p2
	}

	return game{b: p.NextBoard(g.b), current: OtherPlayerID(g.current), p1: g.p1, p2: g.p2}
}

func (g game) IsOver() (bool, PlayerID) {
	// Get all the lines
	lines := AllLinesOf(g.b)

	hasFreeSpace := false
	for _, l := range lines {
		cells := AllCellsOf(l)
		allMatch := true // until proved otherwise
		toMatch := NoPlayer
		for i, c := range cells {
			occ := c.Occupant()

			if i == 0 {
				toMatch = occ
			} else if occ != toMatch {
				allMatch = false
			}

			if occ == NoPlayer {
				hasFreeSpace = true
				allMatch = false
			}
		}
		
		if allMatch {
			return true, toMatch
		}
	}
	return !hasFreeSpace, NoPlayer
}
