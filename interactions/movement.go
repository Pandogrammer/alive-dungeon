package interactions

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
)

type Result int

const (
	Success Result = iota
	Error
)

func Movement(creature c.Creature, world w.World, creatures []c.Creature, movement c.Move) Result {
	var position = creature.Position
	var newX int
	var newY int
	switch movement.Direction {
	case c.Right:
		newX = position.X+ 1
		newY = position.Y
	case c.Left:
		newX = position.X - 1
		newY = position.Y
	case c.Up:
		newX = position.X
		newY = position.Y - 1
	case c.Down:
		newX = position.X
		newY = position.Y + 1
	}

	var newPos = c.Position{newX, newY}

	if collision(newPos, world, creatures) {
		return Error
	}
	return Success
}



func collision(position c.Position, world w.World, creatures []c.Creature) bool {
	if world.OutOfBounds(w.Position(position)) {
		return true
	}
	if world.CellAt(position.X, position.Y) == w.Wall {
		return true
	}
	for _, c := range creatures {
		if c.Position == position {
			return true
		}
	}

	return false
}
