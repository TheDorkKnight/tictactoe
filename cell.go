package main

type Cell interface {
	Occupant() PlayerID
	SetOccupant(pid PlayerID)
}

type cell struct {
	pid PlayerID
}

func (c *cell) Occupant() PlayerID {
	if c == nil {
		return NoPlayer
	}
	return c.pid
}

func (c *cell) SetOccupant(pid PlayerID) {
	if c != nil {
		c.pid = pid
	}
}
