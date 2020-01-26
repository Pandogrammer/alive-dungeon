package world

type Cell int

const (
	Empty Cell = iota
	Wall
)

type World struct {
	Cells [][]Cell
}

type Position struct {
	X int
	Y int
}

func (w World) OutOfBounds(position Position) bool {
	if position.X < 0 || position.X == w.Width() || position.Y < 0 || position.Y == w.Height() {
		return true
	}
	return false
}

func (w World) Height() int {
	return len(w.Cells)
}

func (w World) Width() int {
	return len(w.Cells[0])
}

func (w World) CellAt(x int, y int) Cell {
	return w.Cells[y][x]
}

