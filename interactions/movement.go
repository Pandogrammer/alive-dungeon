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
	var newPos = c.Position{}
	switch movement.Direction {
	case c.Right:
		newPos = c.Position{position.X + 1, position.Y}
	case c.Left:
		newPos = c.Position{position.X - 1, position.Y}
	case c.Up:
		newPos = c.Position{position.X, position.Y - 1}
	case c.Down:
		newPos = c.Position{position.X, position.Y + 1}
	}

	if collision(newPos, world, creatures) {
		return Error
	}
	return Success
}

func collision(position c.Position, world w.World, creatures []c.Creature) bool {
	if position.X < 0 || position.X == len(world.Cells[0]) {
		return true
	}
	if position.Y < 0 || position.Y == len(world.Cells) {
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
