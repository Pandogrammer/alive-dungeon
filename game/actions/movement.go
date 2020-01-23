package actions

import (
	g "alive-dungeon/game"
	w "alive-dungeon/world"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type MoveEvent struct {
	Direction       Direction
	SpawnedCreature g.SpawnedCreature
}

func Move(event MoveEvent) g.SpawnedCreature {
	switch event.Direction {
	case Up :
		return moveUp(event)
	case Down :
		return moveDown(event)
	case Left :
		return moveLeft(event)
	case Right :
		return moveRight(event)
	}
	return event.SpawnedCreature
}

func moveRight(request MoveEvent) g.SpawnedCreature {
	position := request.SpawnedCreature.Position
	newPosition := w.Position{X: position.X + 1, Y: position.Y}
	movedCreature := g.SpawnedCreature{Creature: request.SpawnedCreature.Creature, Position: newPosition}
	return movedCreature
	
}

func moveLeft(request MoveEvent) g.SpawnedCreature {
	position := request.SpawnedCreature.Position
	newPosition := w.Position{X: position.X - 1, Y: position.Y}
	movedCreature := g.SpawnedCreature{Creature: request.SpawnedCreature.Creature, Position: newPosition}
	return movedCreature
	
}

func moveDown(request MoveEvent) g.SpawnedCreature {
	position := request.SpawnedCreature.Position
	newPosition := w.Position{X: position.X, Y: position.Y + 1}
	movedCreature := g.SpawnedCreature{Creature: request.SpawnedCreature.Creature, Position: newPosition}
	return movedCreature
	
}

func moveUp(request MoveEvent) g.SpawnedCreature {
	position := request.SpawnedCreature.Position
	newPosition := w.Position{X: position.X, Y: position.Y - 1}
	movedCreature := g.SpawnedCreature{Creature: request.SpawnedCreature.Creature, Position: newPosition}
	return movedCreature
}
