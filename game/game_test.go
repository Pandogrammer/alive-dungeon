package game

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	world := w.Create(w.CreationRequest{Width: 3, Height: 3})

	game := Create(world)

	assert.Equal(t, game.World, world)
}

func TestAddCreature(t *testing.T) {
	world := w.Create(w.CreationRequest{Width: 3, Height: 3})
	creature := c.Creature{}
	game := Create(world)

	game = game.AddCreature(creature, w.Position{X: 2, Y: 3})

	assert.Equal(t, game.Creatures[0], creature)
}
