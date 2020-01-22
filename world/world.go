package world

type Position struct {
	X int
	Y int
}

type CreationRequest struct {
	Width  int
	Height int
	Walls  []Position
}

type Cell interface{}
type Wall Cell
type Empty Cell

type World struct {
	Cells    [][]Cell
}

func (w World) AddWall(position Position) {
	w.Cells[position.Y][position.X] = new(Wall)
}

func Create(request CreationRequest) World {
	cells := make([][]Cell, request.Width)
	for i := 0; i < request.Width; i++ {
		cells[i] = make([]Cell, request.Height)
		for j := range cells[i] {
			cells[i][j] = new(Empty)
		}
	}
	cells = generateWalls(cells, request)
	return World{Cells: cells}
}

func generateWalls(cells [][]Cell, request CreationRequest) [][]Cell {
	if !(len(request.Walls) > 0) {
		return cells
	}

	for _, w := range request.Walls {
		cells[w.X][w.Y] = new(Wall)
	}

	return cells
}
