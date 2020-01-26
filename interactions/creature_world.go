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

func Movement(world w.World, creature c.Creature, movement c.Move) Result {
	var position = creature.Position
	switch movement.Direction {
	case c.Right:
		var newPos = c.Position{position.X + 1, position.Y}
		if collision(newPos, world){
			return Error
		}
	case c.Left:
		var newPos = c.Position{position.X - 1, position.Y}
		if collision(newPos, world){
			return Error
		}
	case c.Up:
		var newPos = c.Position{position.X, position.Y - 1}
		if collision(newPos, world){
			return Error
		}
	case c.Down:
		var newPos = c.Position{position.X, position.Y + 1}
		if collision(newPos, world){
			return Error
		}
	}
	return Success
}

func collision(position c.Position, world w.World) bool {
	if position.X < 0 || position.X == len(world.Cells[0]){
		return true
	}
	if position.Y < 0 || position.Y == len(world.Cells){
		return true
	}
	if world.CellAt(position.X, position.Y) == w.Wall {
		return true
	}

	return false
}

