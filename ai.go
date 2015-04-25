package main

import (
	"crypto/rand"
)

type AI struct {
	self PlayerID
}

const (
	firstClaimOnLineFitness = 1
	threatenWinFitness = (firstClaimOnLineFitness * 8) + 1
	winFitness = (threatenWinFitness * 8) + 1
)

func (ai AI) NextBoard(b Board) Board {
	// Make an opponent AI to test our ideas against
	opponent := AI{self: OtherPlayerID(ai.self)}

	// Make all possible next boards, 
	// and find the best one
	var bestBoard Board
	var bestFitnessDifferential int
	for i, row := range RowsOf(b) {
		for j, c := range AllCellsOf(row) {
			if c.Occupant() == NoPlayer {
				newB := b.Clone()
				newB.Row(i).Cell(j).SetOccupant(ai.self)
				
				myF := boardFitness(newB, ai.self)
				if myF > winFitness {
					// we can just use this one, no
					// questions asked
					return newB
				}

				opponentB := opponent.NextBoard(newB)
				newFD := myF - boardFitness(opponentB, opponent.self)

				if bestBoard == nil || newFD > bestFitnessDifferential {
					bestBoard = newB
					bestFitnessDifferential = newFD
				} else if newFD == bestFitnessDifferential {
					// toss a coin on whether to use this
					// one or not.
					byt := make([]byte, 1)
					rand.Read(byt)
					if (byt[0] & 0x01) != 0x00 {
						bestBoard = newB
					}
				}
			}
		}
	}
	return bestBoard
}

var p1BoardFitnessMap map[string]int = make(map[string]int)
var p2BoardFitnessMap map[string]int = make(map[string]int)

func boardFitness(b Board, self PlayerID) int {
	if b == nil {
		return 0
	}
	strRep := b.String()

	var fitnessMap map[string]int
	if self == PlayerOne {
		fitnessMap = p1BoardFitnessMap
	} else if self == PlayerTwo {
		fitnessMap = p2BoardFitnessMap
	}

	if fitnessMap != nil {
		f, ok := fitnessMap[strRep]
		if ok {
			return f
		}
	}

	// Get all the lines
	lines := AllLinesOf(b)

	// Evaluate fitness for all lines
	f := 0
	for _, l := range lines {
		f += lineFitness(l, self)
	}

	if fitnessMap != nil {
		fitnessMap[strRep] = f
	}

	return f
}

func lineFitness(l Line, self PlayerID) int {
	myCount := 0
	opponentCount := 0
	cells := AllCellsOf(l)
	for _, c := range cells {
		occupant := c.Occupant()
		if occupant != NoPlayer {
			if occupant == self {
				myCount++
			} else {
				opponentCount++
			}
		}
	}

	if myCount > 0 && opponentCount > 0 {
		// This line isn't worth anything to anyone
		return 0
	}
	return markerFitness(myCount)
}

func markerFitness(numMarkersOnLine int) int {
	switch numMarkersOnLine {
	case 3:
		return winFitness
	case 2:
		return threatenWinFitness
	case 1:
		return firstClaimOnLineFitness
	}
	return 0
}
