package world

type Position struct {
	x int
	y int
}

type CreationRequest struct {
	width  int
	height int
	walls  []Position
}

type Cell interface{}
type Wall Cell

type World struct {
	cells [][]Cell
}

func Create(request CreationRequest) World {
	cells := make([][]Cell, request.width)
	for i := 0; i < request.width; i++ {
		cells[i] = make([]Cell, request.height)
	}
	cells = generateWalls(cells, request)
	return World{cells: cells}
}

func generateWalls(cells [][]Cell, request CreationRequest) [][]Cell {
	if !(len(request.walls) > 0) {
		return cells
	}

	for _, w := range request.walls {
		cells[w.x][w.y] = new(Wall)
	}

	return cells
}
