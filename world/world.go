package world

type Cell int

const (
	Empty Cell = iota
	Wall
)

type World struct {
	Cells [][]Cell
}

func (w World) CellAt(x int, y int) Cell {
	return w.Cells[y][x]
}

type Position struct {
	X int
	Y int
}
