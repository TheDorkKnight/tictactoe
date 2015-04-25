package main

import (
	"fmt"
)

type Human struct {
	self PlayerID
}

func (h Human) NextBoard(b Board) Board {
	str := "Enter the number on the space to place your marker:\n"

	fmt.Print(str)

	gotIt := false
	var cellID int
	for !gotIt {
		n, err := fmt.Scan(&cellID)
		if n < 1 {
			fmt.Println("You didn't enter a number")
		} else if n > 1 {
			fmt.Println("You entered too many numbers")
		} else if err != nil {
			fmt.Println("There was a problem: ", err)
		} else {
			if cellID < 1 {
				fmt.Println("The number you entered was too low: ", cellID)
			} else if cellID > 9 {
				fmt.Println("The number you entered was too high: ", cellID)
			} else if spaceIsOccupied(b, cellID) {
				fmt.Println("The space you entered was already occupied: ", cellID)
			} else {
				gotIt = true
			}
		}
		if !gotIt {
			fmt.Println("Please try again")
		}
	}

	newB := b.Clone()
	r, c := rowAndCol(cellID)
	newB.Row(r).Cell(c).SetOccupant(h.self)
	return newB
}

func rowAndCol(cellID int) (row, col int) {
	return (cellID - 1) / 3, (cellID - 1) % 3
}

func spaceIsOccupied(b Board, cellID int) bool {
	if b == nil || cellID < 1 || cellID > 9 {
		return false
	}

	r, c := rowAndCol(cellID)
	return (b.Row(r).Cell(c).Occupant() != NoPlayer)
}
