package interactions

import (
	c "alive-dungeon/source/creature"
	w "alive-dungeon/source/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovement(t *testing.T) {
	var creature = c.New(c.Create{Position: c.Position{X: 5, Y: 5}})
	var world = w.New(w.Create{Width: 10, Height: 10})
	world = world.Update(w.Modify{Position: w.Position{X: 6, Y: 5}, NewType: w.Wall})
	world = world.Update(w.Modify{Position: w.Position{X: 5, Y: 6}, NewType: w.Wall})

	movement := []struct {
		world    w.World
		creature c.Creature
		move     c.Move
		result   Result
	}{
		{world: world, creature: creature, move: c.Move{Direction: c.Right}, result: Error},
		{world: world, creature: creature, move: c.Move{Direction: c.Left}, result: Success},
		{world: world, creature: creature, move: c.Move{Direction: c.Up}, result: Success},
		{world: world, creature: creature, move: c.Move{Direction: c.Down}, result: Error},
	}

	for _, test := range movement {
		var result = Movement(test.creature, test.world, nil, test.move)

		assert.Equal(t, test.result, result)
	}
}
