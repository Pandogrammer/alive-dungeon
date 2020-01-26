package creature

func (c Creature) Update(message interface{}) Creature {
	switch message.(type) {
	case Move:
		return c.move(message.(Move))
	}
	return c
}

func New(message Create) Creature {
	return Creature{message.Position}
}

func (c Creature) move(move Move) Creature {
	var newPosition = Position{}
	switch move.Direction {
	case Right:
		newPosition = Position{c.Position.X + 1, c.Position.Y}
	case Left:
		newPosition = Position{c.Position.X - 1, c.Position.Y}
	case Down:
		newPosition = Position{c.Position.X, c.Position.Y + 1}
	case Up:
		newPosition = Position{c.Position.X, c.Position.Y - 1}
	}

	c.Position = newPosition
	return c
}
