package world

func New(message Create) World {
	return create(message.Height, message.Width)
}

func (w World) Update(message interface{}) World {
	switch message.(type) {
	case Modify:
		return w.modify(message.(Modify))
	}
	return w
}

func (w World) modify(message Modify) World {
	var position = message.Position
	w.Cells[position.Y][position.X] = message.NewType
	return w
}

func create(height int, width int) World {
	cells := make([][]Cell, width)
	for i := 0; i < width; i++ {
		cells[i] = make([]Cell, height)
		for j := range cells[i] {
			cells[i][j] = Empty
		}
	}
	return World{cells}
}
