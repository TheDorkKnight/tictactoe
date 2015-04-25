package main

import (
	"strings"
)

const (
	TopRow = 0
	BottomRow = 2
)

const (
	LeftColumn = 0
	RightColumn = 2
)

const (
	TL2BRDiagonal = 0
	BL2TRDiagonal = 1
)

type Board interface {
	Row(rid int) Line
	Column(cid int) Line
	Diagonal(did int) Line
	PrettyString() string
	String() string
	Clone() Board
}

type board struct {
	cells [3][3]cell
}

func (b *board) Row(rid int) Line {
	if b == nil {
		return nil
	}

	if rid < TopRow || rid > BottomRow {
		return nil
	}
	return makeLine(&(b.cells[rid][0]), &b.cells[rid][1], &b.cells[rid][2])
}

func (b *board) Column(cid int) Line {
	if b == nil {
		return nil
	}

	if cid < LeftColumn || cid > RightColumn {
		return nil
	}
	return makeLine(&b.cells[0][cid], &b.cells[1][cid], &b.cells[2][cid])
}

func (b *board) Diagonal(did int) Line {
	if b == nil {
		return nil
	}

	switch (did) {
	case TL2BRDiagonal:
		return makeLine(&b.cells[0][0], &b.cells[1][1], &b.cells[2][2])
	case BL2TRDiagonal:
		return makeLine(&b.cells[2][0], &b.cells[1][1], &b.cells[0][2])
	}
	return nil
}

const (
	underscoresTop    = " _____ _____ _____\n"
	underscoresNormal = "|_____|_____|_____|\n"
	pipes             = "|     |     |     |\n"
	xTop        = " \\ / |"
	xMid        =  "  X  |"
	xBot        = " / \\ |"
	oTop        =  "  _  |"
	oMid        =  " (_) |"
	oBot        =  "     |"
	unTopOrBot  = oBot
)

var (
	nums        = [...]string{"1","2","3","4","5","6","7","8","9"}
)

func (b *board) PrettyString() string {
	if b == nil {
		return ""
	}

	var strs []string
	strs = append(strs, underscoresTop)
	for i := 0; i < 3; i++ {
		row := b.Row(i)

		strs = append(strs, pipes)

		strs = append(strs, "|")
		for j := 0; j < 3; j++ {
			switch row.Cell(j).Occupant() {
			case PlayerOne:
				strs = append(strs, xTop)
			case PlayerTwo:
				strs = append(strs, oTop)
			default:
				strs = append(strs, unTopOrBot)
			}
		}
		strs = append(strs, "\n")

		strs = append(strs, "|")
		for j := 0; j < 3; j++ {
			switch row.Cell(j).Occupant() {
			case PlayerOne:
				strs = append(strs, xMid)
			case PlayerTwo:
				strs = append(strs, oMid)
			default:
				strs = append(strs, "  ")
				strs = append(strs, nums[(3*i) + j])
				strs = append(strs, "  |")
			}
		}
		strs = append(strs, "\n")

		strs = append(strs, "|")
		for j := 0; j < 3; j++ {
			switch row.Cell(j).Occupant() {
			case PlayerOne:
				strs = append(strs, xBot)
			case PlayerTwo:
				strs = append(strs, oBot)
			default:
				strs = append(strs, unTopOrBot)
			}
		}
		strs = append(strs, "\n")
		strs = append(strs, underscoresNormal)
	}
	return strings.Join(strs, "")
}

func (b *board) String() string {
	if b == nil {
		return ""
	}

	var strs []string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			strs = append(strs, b.cells[i][j].pid.String())
		}
		strs = append(strs, "\n")
	}
	return strings.Join(strs, "")
}

func (b *board) Clone() Board {
	var clone board = *b
	return &clone
}


func RowsOf(b Board) []Line {
	if b == nil {
		return nil
	}
	lines := make([]Line, 3)
	for i := 0; i < 3; i++ {
		lines[i] = b.Row(i)
	}
	return lines
}

func AllLinesOf(b Board) []Line {
	if b == nil {
		return nil
	}
	lines := make([]Line, 8)
	for i := 0; i < 3; i++ {
		lines[i] = b.Row(i)
		lines[3+i] = b.Column(i)
		if i < 2 {
			lines[6+i] = b.Diagonal(i)
		}
	}
	return lines
}

func DiagonalsOf(b Board) []Line {
	if b == nil {
		return nil
	}
	lines := make([]Line, 2)
	lines[0] = b.Diagonal(0)
	lines[1] = b.Diagonal(1)
	return lines
}
