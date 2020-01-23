package game

import (
	c "alive-dungeon/creature"
	w "alive-dungeon/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	world := w.Create(w.CreationEvent{Width: 3, Height: 3})

	game := Create(world)

	assert.Equal(t, game.World, world)
}

func TestAddCreature(t *testing.T) {
	world := w.Create(w.CreationEvent{Width: 3, Height: 3})
	creature := SpawnedCreature{c.Creature{}, w.Position{X: 2, Y: 3}}
	game := Create(world)

	game = game.AddCreature(creature)

	assert.Equal(t, game.Creatures[0], creature)
}


