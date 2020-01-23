package actions

import (
	c "alive-dungeon/creature"
	g "alive-dungeon/game"
	w "alive-dungeon/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

var game = g.Game{}
var world = w.World{}
var creature = c.Creature{}

func init() {
	request := w.CreationEvent{Width: 3, Height: 3, Walls: []w.Position{{1, 1}}}
	world = w.Create(request)
	game = g.Create(world)
}

func TestMovement(t *testing.T) {
	//given a creature and a movement request
	spawnedCreature := g.SpawnedCreature{Creature: creature, Position: w.Position{X: 2, Y: 2}}
	movement := []struct {
		request     MoveEvent
		newPosition w.Position
	}{
		{request: MoveEvent{SpawnedCreature: spawnedCreature, Direction: Up},
			newPosition: w.Position{X: 2, Y: 1}},
		{request: MoveEvent{SpawnedCreature: spawnedCreature, Direction: Down},
			newPosition: w.Position{X: 2, Y: 3}},
		{request: MoveEvent{SpawnedCreature: spawnedCreature, Direction: Left},
			newPosition: w.Position{X: 1, Y: 2}},
		{request: MoveEvent{SpawnedCreature: spawnedCreature, Direction: Right},
			newPosition: w.Position{X: 3, Y: 2}},
	}

	for _, move := range movement {
		//when move
		spawnedCreature = Move(move.request)

		//then creature is moved
		assert.Equal(t, move.newPosition, spawnedCreature.Position)
	}

}
