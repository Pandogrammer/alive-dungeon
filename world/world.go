package world

type Position struct {
	X int
	Y int
}

type CreationEvent struct {
	Width  int
	Height int
	Walls  []Position
}

type Cell int

const (
	Empty Cell = iota
	Wall
)

type World struct {
	Cells [][]Cell
}

func (w World) AddWall(position Position) {
	w.Cells[position.Y][position.X] = Wall
}

func Create(request CreationEvent) World {
	cells := make([][]Cell, request.Width)
	for i := 0; i < request.Width; i++ {
		cells[i] = make([]Cell, request.Height)
		for j := range cells[i] {
			cells[i][j] = Empty
		}
	}
	cells = generateWalls(cells, request)
	return World{Cells: cells}
}

func generateWalls(cells [][]Cell, request CreationEvent) [][]Cell {
	if !(len(request.Walls) > 0) {
		return cells
	}

	for _, w := range request.Walls {
		cells[w.Y][w.X] = Wall
	}

	return cells
}
