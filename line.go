package main

const (
	firstCell = 0
	thirdCell = 2
)

type Line interface {
	Cell(cid int) Cell
}

type line struct {
	cells [3]Cell
}

func (l line) Cell(cid int) Cell {
	if cid < firstCell || cid > thirdCell {
		return nil
	}
	return l.cells[cid]
}

func makeLine(first, second, third Cell) Line {
	return line{cells: [3]Cell{first, second, third}}
}

func AllCellsOf(l Line) []Cell {
	if l == nil {
		return nil
	}
	cells := make([]Cell, 3)
	for i := 0; i < 3; i++ {
		cells[i] = l.Cell(i)
	}
	return cells
}
